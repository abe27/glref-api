package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func GlrefController(c *fiber.Ctx) error {
	var r models.Response
	var gl []models.Glref
	if err := configs.StoreFormula.
		Preload("BOOK").
		Preload("BRANCH").
		Preload("COOR").
		Preload("CORP").
		Preload("CORRECTB").
		Preload("DEPT").
		Preload("JOB").
		Preload("PROJ").
		Preload("TOWHOUSE").
		Preload("FROMWHOUSE").
		Where("FCSKID", c.Query("fcskid")).Find(&gl).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Data = &gl
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlrefPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.GlrefForm
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	// Get EmpID
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	uid, err := services.ValidateToken(token)
	if err != nil {
		r.Message = "Token is Expired"
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}

	// Form Insert
	db := configs.StoreFormula
	// Begin transaction
	tx := db.Begin()
	// Commit the transaction in case success
	onYear, _ := strconv.Atoi(time.Now().Format("2006"))
	thYear := fmt.Sprintf("%d", (onYear + 543))
	var rnn int64
	if err := tx.Select("FCCODE").Where("FCCODE LIKE ?", (thYear + ((time.Now().Format("20060102"))[4:6]))[2:6]+"%").Model(&models.Glref{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	var book models.Booking
	if err := tx.Preload("REFTYPE").Where("FCSKID", frm.Booking).First(&book).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var sect models.Section
	if err := tx.Preload("CORP").Preload("DEPT").Where("FCDEPT", frm.Department).First(&sect).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var fcamt float64 = 0
	for _, i := range frm.Items {
		fcamt += i.Qty
	}

	empID := fmt.Sprintf("%s", uid)
	fccode := fmt.Sprintf("%s%04d", (thYear + ((time.Now().Format("20060102"))[4:6]))[2:6], (rnn + 1))
	var glref models.Glref
	glref.FCCODE = fccode
	glref.FCREFNO = fmt.Sprintf("%s%s", strings.Trim(book.FCPREFIX, " "), fccode)
	glref.FCREFTYPE = book.FCREFTYPE
	glref.FCRFTYPE = book.REFTYPE.FCRFTYPE
	glref.FCSTEP = frm.Step
	glref.FDDATE = frm.RecDate
	glref.FCBRANCH = frm.Branch
	glref.FNAMT = fcamt
	glref.FCPROJ = frm.Proj
	glref.FCJOB = frm.Job
	glref.FCCOOR = frm.Coor
	glref.FCCORP = frm.Corp
	glref.FCDEPT = frm.Department
	glref.FCSECT = sect.FCSKID
	glref.FCBOOK = book.FCSKID
	glref.FCCORRECTB = empID
	glref.FMMEMDATA = strings.ToUpper(frm.InvoiceNo)
	glref.FCTOWHOUSE = frm.Whs

	if err := tx.Create(&glref).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var seq int64 = 1
	for _, i := range frm.Items {
		var prod models.Product
		if err := tx.Where("FCSKID", i.Product).First(&prod).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		var refProd models.Refprod
		refProd.FCSEQ = fmt.Sprintf("%03d", seq)
		refProd.FCGLREF = glref.FCSKID
		refProd.FDDATE = glref.FDDATE
		refProd.FCIOTYPE = glref.FCSTEP
		refProd.FCRFTYPE = book.REFTYPE.FCRFTYPE
		refProd.FCREFTYPE = book.FCREFTYPE
		refProd.FCPRODTYPE = book.REFTYPE.FCRFTYPE
		refProd.FCCORP = frm.Corp
		refProd.FCBRANCH = frm.Branch
		refProd.FCWHOUSE = frm.Whs
		refProd.FCJOB = frm.Job
		refProd.FCSECT = sect.FCSKID
		refProd.FCDEPT = sect.FCDEPT
		refProd.FCPROD = i.Product
		refProd.FCPRODTYPE = prod.FCPRTYPE
		refProd.FNUMQTY = 1
		refProd.FNQTY = i.Qty
		refProd.FNPRICE = i.Price
		refProd.FCUM = i.Unit
		refProd.FCUMSTD = i.Unit
		refProd.FCSTUM = i.Unit
		refProd.FCSTUMSTD = i.Unit
		refProd.FNSTUMQTY = 1
		refProd.FTDATETIME = time.Now()
		refProd.FTLASTEDIT = time.Now()
		refProd.FTLASTUPD = time.Now()

		if err := tx.Create(&refProd).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		var stock models.Stock
		tx.First(&stock, &models.Stock{FCPROD: prod.FCSKID, FCWHOUSE: frm.Whs})
		stock.FCCORP = frm.Corp
		stock.FCBRANCH = frm.Branch
		stock.FCWHOUSE = frm.Whs
		stock.FCPROD = prod.FCSKID
		stock.FDDATE = glref.FDDATE
		switch frm.Step {
		case "I":
			stock.FNQTY = stock.FNQTY + i.Qty
		default:
			if stock.FNQTY > 0 {
				stock.FNQTY = stock.FNQTY - i.Qty
			}
		}

		if stock.FCSKID == "" {
			stock.FTDATETIME = time.Now()
		}
		stock.FTLASTUPD = time.Now()
		stock.FTLASTEDIT = time.Now()
		if err := tx.Save(&stock).Error; err != nil {
			tx.Rollback()
			r.Message = fmt.Sprintf("Failed transection on Stock: %v", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		seq++
	}

	// Glref History
	var glrefHistory models.GlrefHistory
	glrefHistory.FCSKID = glref.FCSKID
	glrefHistory.FCCODE = glref.FCCODE
	glrefHistory.FCREFNO = glref.FCREFNO
	glrefHistory.FCDATE = glref.FDDATE
	glrefHistory.FCINVOICE = strings.ToUpper(frm.InvoiceNo)
	if err := configs.Store.Create(&glrefHistory).Error; err != nil {
		tx.Rollback()
		r.Message = fmt.Sprintf("Failed create glref history: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	tx.Commit()
	// End
	r.Message = fmt.Sprintf("%s <> %s", fccode, uid)

	// msg := fmt.Sprintf("\nบันทึก%s\nเลขที่: %s \nสินค้า: %d รายการ\nจำนวน: %d\nเรียบร้อยแล้ว\n%s", book.FCNAME, glref.FCREFNO, len(frm.REFPROD), int(fcamt), time.Now().Format("2006-01-02 15:04:05"))
	// var line models.Linenotify
	// if err := configs.Store.First(&line, &models.Linenotify{Jobs: book.FCREFTYPE}).Error; err == nil {
	// 	if line.Token != "" {
	// 		go services.LineNotify(line.Token, msg)
	// 	}
	// }

	msg := fmt.Sprintf("\nบันทึก%s\nเลขที่: %s \nสินค้า: %d รายการ\nจำนวน: %d\nเรียบร้อยแล้ว\n%s", book.FCNAME, glref.FCREFNO, len(frm.Items), int(fcamt), time.Now().Format("2006-01-02 15:04:05"))
	go services.LineNotify(configs.APP_LINE_TOKEN, msg)

	r.Data = &glref
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func GlrefTransferController(c *fiber.Ctx) error {
	db := configs.StoreFormula
	var r models.Response
	r.Message = "Patch"

	// Get EmpID
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	uid, err := services.ValidateToken(token)
	if err != nil {
		r.Message = "Token is Expired"
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}

	if c.Query("fcskid") == "" {
		r.Message = "GLREF FCSKID Required!"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if c.Query("pono") == "" {
		r.Message = "PONO Required!"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	tx := db.Begin()
	store := configs.Store.Begin()
	var glRef models.Glref
	if err := tx.First(&glRef, &models.Glref{FCSKID: c.Query("fcskid")}).Error; err != nil {
		r.Message = "ไม่พบข้อมูลที่ระบุ"
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Update History
	var glHistory models.GlrefHistory
	if err := store.First(&glHistory, &models.GlrefHistory{FCSKID: glRef.FCSKID}).Error; err != nil {
		tx.Rollback()
		r.Message = "ไม่พบข้อมูลที่ระบุ"
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var orderH models.Orderh
	if err := tx.First(&orderH, &models.Orderh{FCREFNO: strings.ToUpper(c.Query("pono"))}).Error; err != nil {
		r.Message = "ไม่พบข้อมูลเลขที่ PO นี้"
		glHistory.FCREMARK = fmt.Sprintf("ไม่พบข้อมูลเลขที่ %s นี้", strings.ToUpper(c.Query("pono")))
		glHistory.FCSTATUS = 2
		if err := store.Save(&glHistory).Error; err != nil {
			tx.Rollback()
			store.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		store.Commit()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var refProd []models.Refprod
	if err := tx.
		Preload("Corp").
		Preload("Branch").
		Preload("Dept").
		Preload("Sect").
		Preload("Job").
		Preload("Glhead").
		Preload("Glref").
		Preload("Prod").
		Preload("Unit").
		Preload("UnitSTD").
		Preload("Stum").
		Preload("StumStd").
		Preload("WHouse").
		Preload("Proj").
		Preload("Gl").
		Find(&refProd, &models.Refprod{FCGLREF: c.Query("fcskid")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Check Part from orderi
	var listOrderI []models.OrderiView
	for _, p := range refProd {
		var prodOrderI models.OrderiView
		if err := tx.Where("FCSTEP", "1").First(&prodOrderI, &models.OrderiView{FCORDERH: orderH.FCSKID, FCPROD: p.FCPROD}).Error; err != nil {
			r.Message = "พบสินค้าไม่ตรงกับเอกสาร"
			glHistory.FCREMARK = "พบสินค้าไม่ตรงกับเอกสาร"
			glHistory.FCSTATUS = 2
			if err := store.Save(&glHistory).Error; err != nil {
				tx.Rollback()
				store.Rollback()
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
			store.Commit()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		if p.FNQTY > prodOrderI.FNQTY {
			r.Message = "ระบุจำนวนเกินกับเอกสาร"
			glHistory.FCREMARK = "ระบุจำนวนเกินกับเอกสาร"
			glHistory.FCSTATUS = 2
			if err := store.Save(&glHistory).Error; err != nil {
				tx.Rollback()
				store.Rollback()
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
			store.Commit()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		prodOrderI.FNRECEIVEQTY = p.FNQTY
		listOrderI = append(listOrderI, prodOrderI)
	}

	// CREATE GLHEAD
	var glhead models.Glhead
	glhead.FCCORP = glRef.FCCORP
	glhead.FCBRANCH = glRef.FCBRANCH
	glhead.FDDATE = glRef.FDDATE
	var accBook models.Accbook
	if err := tx.First(&accBook, &models.Accbook{FCCODE: "PD"}).Error; err != nil {
		r.Message = "ไม่พบข้อมูล ACCBOOK!"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	glhead.FCACCBOOK = accBook.FCSKID

	onYear, _ := strconv.Atoi(time.Now().Format("2006"))
	thYear := fmt.Sprintf("%d", (onYear + 543))
	var rnn int64
	if err := tx.Select("FCCODE").Where("FCCODE LIKE ?", (thYear + ((time.Now().Format("20060102"))[4:6]))[2:6]+"%").Model(&models.Glhead{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	glhead.FCCODE = fmt.Sprintf("%s%04d", (thYear + ((time.Now().Format("20060102"))[4:6]))[2:6], (rnn + 1))
	glhead.FCCORRECTB = fmt.Sprintf("%s", uid)
	if err := tx.Create(&glhead).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// fmt.Println("GLHEAD: %s", glhead.FCSKID)
	// UPDATE GLREF
	var book models.Booking
	if err := tx.First(&book, &models.Booking{FCCODE: "003", FCREFTYPE: "BI"}).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	glRef.FCBOOK = book.FCSKID
	glRef.FCGLHEAD = glhead.FCSKID
	glRef.FCRFTYPE = "B"
	glRef.FCREFTYPE = "BI"
	glRef.FCSTEP = "P"
	glRef.FTLASTEDIT = time.Now()
	glRef.FTLASTUPD = time.Now()
	if err := tx.Save(&glRef).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// var sumRefProd float64 = 0
	// var sumCtn float64 = 0
	for _, i := range listOrderI {
		// fmt.Println("GLREF: %s PROD: %s ORDERH: %s ORDERI: %s", glRef.FCSKID, i.FCPROD, i.FCORDERH, i.FCSKID)
		var refProd models.Refprod
		if err := tx.Select("FCSKID,FNQTY,FCSEQ,FNPRICE").First(&refProd, &models.Refprod{FCGLREF: glRef.FCSKID, FCPROD: i.FCPROD}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// UPDATE REFPROD
		// sumRefProd += refProd.FNQTY
		if err := tx.Model(&refProd).Updates(&models.Refprod{
			FCGLHEAD:   glhead.FCSKID,
			FCRFTYPE:   "B",
			FCREFTYPE:  "BI",
			FTLASTEDIT: time.Now(),
			FTLASTUPD:  time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		backQty := i.FNBACKQTY - i.FNRECEIVEQTY
		orderIStatus := "1"
		if backQty == 0 {
			orderIStatus = "P"
		}

		// UPDATE ORDERI
		if err := tx.Model(&models.Orderi{FCSKID: i.FCSKID}).Updates(&models.Orderi{
			FCSTEP:     orderIStatus,
			FNBACKQTY:  backQty,
			FNPRICE:    refProd.FNPRICE,
			FTLASTEDIT: time.Now(),
			FTLASTUPD:  time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		// fmt.Println("orderIStatus: ", orderIStatus)
		// sumCtn += orderIQty

		// UPDATE GLREF HISTORY
		var orderH models.Orderh
		if err := tx.Select("FCREFNO").First(&orderH, &models.Orderh{FCSKID: listOrderI[0].FCORDERH}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		var prod models.Product
		if err := tx.Select("FCSKID,FCACCSCRED").First(&prod, &models.Product{FCSKID: i.FCPROD}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		var acChart models.Acchart
		if err := tx.First(&acChart, &models.Acchart{FCSKID: prod.FCACCBCRED}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		// CREATE GL Round 1
		var glFirst models.Gl
		glFirst.FCACCHART = acChart.FCSKID
		glFirst.FCBRANCH = glRef.FCBRANCH
		glFirst.FCCORP = glRef.FCCORP
		glFirst.FCSECT = glRef.FCSECT
		glFirst.FCDEPT = glRef.FCDEPT
		glFirst.FCGLHEAD = glRef.FCGLHEAD
		glFirst.FCSEQ = fmt.Sprintf("%04d", 1)
		glFirst.FDDATE = glRef.FDDATE
		glFirst.FNAMT = refProd.FNPRICE
		if err := tx.Create(&glFirst).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// CREATE GL "1061", "2021"
		var accChart []models.Acchart
		if err := tx.Where("FCCODE IN ?", [2]string{"1061", "2021"}).Find(&accChart).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		var vatSum float64
		runn := 0
		for _, ac := range accChart {
			var seqGl int64
			if err := tx.Select("FCSKID").Find(&models.Gl{}, &models.Gl{FCGLHEAD: glRef.FCGLHEAD}).Count(&seqGl).Error; err != nil {
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}

			var gl models.Gl
			gl.FCACCHART = ac.FCSKID
			gl.FCBRANCH = glRef.FCBRANCH
			gl.FCCORP = glRef.FCCORP
			gl.FCSECT = glRef.FCSECT
			gl.FCDEPT = glRef.FCDEPT
			gl.FCGLHEAD = glRef.FCGLHEAD
			gl.FCSEQ = fmt.Sprintf("%04d", (seqGl + 1))
			gl.FDDATE = glRef.FDDATE
			vatSum = (refProd.FNPRICE * 0.07)
			if runn == 0 {
				gl.FNAMT = vatSum
			} else {
				gl.FNAMT = 0 - (vatSum + refProd.FNPRICE)
			}

			if err := tx.Create(&gl).Error; err != nil {
				tx.Rollback()
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
			runn++
		}
		// CREATE NOTCUT
		var noteCut models.Notecut
		noteCut.FCBRANCH = glRef.FCBRANCH
		noteCut.FCCHILDH = listOrderI[0].FCORDERH
		noteCut.FCCHILDI = i.FCSKID
		noteCut.FCCORP = glRef.FCCORP
		noteCut.FCMASTERH = glRef.FCSKID
		noteCut.FCMASTERI = refProd.FCSKID
		noteCut.FNQTY = i.FNRECEIVEQTY
		if err := tx.Create(&noteCut).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	var sumCtn int64
	if err := tx.Raw(fmt.Sprintf("select count(FCSKID) from ORDERI where FCORDERH='%s' and FCSTEP='1'", listOrderI[0].FCORDERH)).Scan(&sumCtn).Error; err != nil {
		tx.Rollback()
		store.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	// UPDATE ORDERH
	orderStatus := "P"
	if sumCtn > 0 {
		orderStatus = "1"
	}
	// fmt.Println("ORDERID: ", listOrderI[0].FCORDERH, " SUM: ", sumCtn, "STATUS: ", orderStatus)
	if err := tx.Model(&models.Orderh{FCSKID: listOrderI[0].FCORDERH}).Updates(&models.Orderh{
		FCSTEP:     orderStatus,
		FTLASTEDIT: time.Now(),
		FTLASTUPD:  time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		store.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Update History
	glHistory.FCPONO = strings.ToUpper(c.Query("pono"))
	glHistory.FCSTATUS = 1

	if err := store.Save(&glHistory).Error; err != nil {
		tx.Rollback()
		store.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	tx.Commit()
	store.Commit()
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlrefHistoryController(c *fiber.Ctx) error {
	var r models.Response
	var gl []models.GlrefHistory
	if c.Query("fcskid") != "" {
		if err := configs.Store.Where("fcsk_id", c.Query("fcskid")).Find(&gl).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&gl)
		}
	}

	if err := configs.Store.Limit(150).Order("FCDATE").Find(&gl).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&gl)
	}
	r.Data = &gl
	return c.Status(fiber.StatusOK).JSON(&r)
}
