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
	"github.com/samber/lo"
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
	var fvatamt float64 = 0
	for _, i := range frm.Items {
		fcamt += (i.Price * i.Qty)
		fvatamt += (i.Price * i.Qty) * 0.07
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
	glref.FNVATAMT = fvatamt
	glref.FCPROJ = frm.Proj
	glref.FCJOB = frm.Job
	glref.FCCOOR = frm.Coor
	glref.FCCORP = frm.Corp
	glref.FCDEPT = frm.Department
	glref.FCSECT = sect.FCSKID
	glref.FCBOOK = book.FCSKID
	glref.FCCORRECTB = empID
	glref.FMMEMDATA = strings.ToUpper(frm.InvoiceNo)
	if frm.FromWhs != "" {
		glref.FCFRWHOUSE = frm.FromWhs
	}
	glref.FCTOWHOUSE = frm.ToWhs
	glref.FCCREATEBY = empID
	glref.FCVATCOOR = frm.Coor
	// glref.FCCREATETY = empID
	glref.FCDELICOOR = frm.Coor

	if err := tx.Create(&glref).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	switch frm.Prefix {
	case "PVF1":
		//start refprod
		numRun := 1
		for _, typeName := range [2]string{"O", "I"} {
			strRound := fmt.Sprintf("%03d", numRun)
			fmt.Println(typeName + " ::==> " + strRound)
			var seq int64 = 1
			for _, i := range frm.Items {
				var prod models.Product
				if err := tx.Where("FCCODE", i.Product).First(&prod).Error; err != nil {
					tx.Rollback()
					r.Message = err.Error()
					return c.Status(fiber.StatusInternalServerError).JSON(&r)
				}

				var refProd models.Refprod
				refProd.FCSEQ = fmt.Sprintf("%03d", seq)
				refProd.FCGLREF = glref.FCSKID
				refProd.FDDATE = glref.FDDATE
				refProd.FCIOTYPE = typeName
				refProd.FCRFTYPE = book.REFTYPE.FCRFTYPE
				refProd.FCREFTYPE = book.FCREFTYPE
				refProd.FCCOOR = frm.Coor
				refProd.FCCORP = frm.Corp
				refProd.FCBRANCH = frm.Branch
				switch typeName {
				case "I":
					refProd.FCWHOUSE = frm.ToWhs
				default:
					refProd.FCWHOUSE = frm.FromWhs
				}

				refProd.FCPROJ = frm.Proj
				refProd.FCJOB = frm.Job
				refProd.FCSECT = sect.FCSKID
				refProd.FCDEPT = sect.FCDEPT
				refProd.FCPROD = prod.FCSKID
				refProd.FCPRODTYPE = prod.FCTYPE
				refProd.FNUMQTY = 1
				refProd.FNQTY = i.Qty
				// refProd.FNPRICE = i.Price
				refProd.FNPRICE = 0
				refProd.FCUM = prod.FCUM
				refProd.FCUMSTD = prod.FCUM
				refProd.FCSTUM = prod.FCUM
				refProd.FCSTUMSTD = prod.FCUM
				refProd.FNSTUMQTY = 1
				// refProd.FNCOSTAMT = i.Price * i.Qty
				// refProd.FNVATAMT = (i.Price * i.Qty) * 0.07
				// refProd.FNPRICEKE = i.Price
				refProd.FNCOSTAMT = 0
				refProd.FNVATAMT = 0
				refProd.FNPRICEKE = 0
				refProd.FCCREATEBY = empID

				if err := tx.Create(&refProd).Error; err != nil {
					tx.Rollback()
					r.Message = err.Error()
					return c.Status(fiber.StatusInternalServerError).JSON(&r)
				}

				var stock models.Stock
				tx.First(&stock, &models.Stock{FCPROD: prod.FCSKID, FCWHOUSE: frm.ToWhs})
				stock.FCCORP = frm.Corp
				stock.FCBRANCH = frm.Branch
				stock.FCWHOUSE = frm.ToWhs
				stock.FCPROD = prod.FCSKID
				stock.FDDATE = glref.FDDATE
				switch typeName {
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
			numRun++
		}

	case "ADJ":
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
			refProd.FCIOTYPE = frm.Step
			refProd.FCRFTYPE = book.REFTYPE.FCRFTYPE
			refProd.FCREFTYPE = book.FCREFTYPE
			refProd.FCCOOR = frm.Coor
			refProd.FCCORP = frm.Corp
			refProd.FCBRANCH = frm.Branch
			refProd.FCWHOUSE = frm.ToWhs
			refProd.FCPROJ = frm.Proj
			refProd.FCJOB = frm.Job
			refProd.FCSECT = sect.FCSKID
			refProd.FCDEPT = sect.FCDEPT
			refProd.FCPROD = i.Product
			refProd.FCPRODTYPE = prod.FCTYPE
			refProd.FNUMQTY = 1
			refProd.FNQTY = i.Qty
			refProd.FNPRICE = i.Price
			refProd.FCUM = i.Unit
			refProd.FCUMSTD = i.Unit
			refProd.FCSTUM = i.Unit
			refProd.FCSTUMSTD = i.Unit
			refProd.FNSTUMQTY = 1
			refProd.FNCOSTAMT = i.Price * i.Qty
			refProd.FNVATAMT = (i.Price * i.Qty) * 0.07
			refProd.FNPRICEKE = i.Price
			refProd.FCCREATEBY = empID

			if err := tx.Create(&refProd).Error; err != nil {
				tx.Rollback()
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}

			var stock models.Stock
			tx.First(&stock, &models.Stock{FCPROD: prod.FCSKID, FCWHOUSE: frm.ToWhs})
			stock.FCCORP = frm.Corp
			stock.FCBRANCH = frm.Branch
			stock.FCWHOUSE = frm.ToWhs
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
	}

	// Glref History
	var glrefHistory models.GlrefHistory
	glrefHistory.FCTYPE = frm.Prefix
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
	// tx.Rollback()
	// End
	r.Message = fmt.Sprintf("%s <> %s", fccode, uid)

	fnQty := 0
	for _, q := range frm.Items {
		fnQty += int(q.Qty)
	}
	msg := fmt.Sprintf("\nบันทึก%s\nเลขที่: %s \nสินค้า: %d รายการ\nจำนวน: %d %s\nเรียบร้อยแล้ว\n%s", book.FCNAME, glref.FCREFNO, len(frm.Items), int(fnQty), "ชิ้น", time.Now().Format("2006-01-02 15:04:05"))
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
	var proArr []string // list of products acchart
	var vatAmt float64 = 0
	var fnAmt float64 = 0      // FNAMTKE
	var fnVatAmt float64 = 0   // FNVATAMTKE
	var fnSumPrice float64 = 0 //
	var listOrderI []models.OrderiView
	for _, p := range refProd {
		vatAmt += p.FNVATAMT
		fnAmt += p.FNCOSTAMT
		fnVatAmt += p.FNPRICEKE
		fnSumPrice += (p.FNCOSTAMT)

		// refProd.FNCOSTAMT = i.Price * i.Qty
		// refProd.FNVATAMT = (i.Price * i.Qty) * 0.07
		// refProd.FNPRICEKE = i.Price
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

		if prodOrderI.FNBACKQTY <= 0 {
			r.Message = "สินค้ารับเข้าคลังหมดไปแล้ว"
			glHistory.FCREMARK = "สินค้ารับเข้าคลังหมดไปแล้ว"
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
		proArr = append(proArr, p.Prod.FCACCBCRED)
	}

	// CREATE GLHEAD
	var glhead models.Glhead
	glhead.FCCORP = glRef.FCCORP
	glhead.FCBRANCH = glRef.FCBRANCH
	glhead.FDDATE = glRef.FDDATE
	glhead.FMREMARK = glRef.FMMEMDATA
	glhead.FCCREATEBY = glRef.FCCREATEBY
	var accBook models.Accbook
	if err := tx.First(&accBook, &models.Accbook{FCCODE: "PD"}).Error; err != nil {
		r.Message = "ไม่พบข้อมูล ACCBOOK!"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	glhead.FCACCBOOK = accBook.FCSKID

	onYear, _ := strconv.Atoi(time.Now().Format("2006"))
	onMonth, _ := strconv.Atoi(time.Now().Format("200601"))
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

	///// CREATE GL Round 1
	// fmt.Println(lo.Uniq(proArr))
	var acChart []models.Acchart
	if err := tx.Select("FCSKID").Where("FCSKID IN ?", lo.Uniq(proArr)).Find(&acChart).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	glRnd := 1
	for _, ac := range acChart {
		var glFirst models.Gl
		glFirst.FCACCHART = ac.FCSKID
		glFirst.FCBRANCH = glRef.FCBRANCH
		glFirst.FCCORP = glRef.FCCORP
		glFirst.FCSECT = glRef.FCSECT
		glFirst.FCDEPT = glRef.FCDEPT
		glFirst.FCGLHEAD = glhead.FCSKID
		glFirst.FCSEQ = fmt.Sprintf("%04d", glRnd)
		glFirst.FDDATE = glRef.FDDATE
		glFirst.FNAMT = fnSumPrice //refProd.FNPRICE
		glFirst.FCJOB = glRef.FCJOB
		glFirst.FCPROJ = glRef.FCPROJ
		glFirst.FCCREATEBY = glRef.FCCREATEBY
		glFirst.FCCORRECTB = glRef.FCCREATEBY
		if err := tx.Create(&glFirst).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		glRnd++
	}

	// CREATE GL "1061", "2021"
	var accChart []models.Acchart
	if err := tx.Where("FCCODE IN ?", [2]string{"1061", "2021"}).Find(&accChart).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	runn := 0
	for _, ac := range accChart {
		var seqGl int64
		if err := tx.Select("FCSKID").Find(&models.Gl{}, &models.Gl{FCGLHEAD: glhead.FCSKID}).Count(&seqGl).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		var gl models.Gl
		gl.FCACCHART = ac.FCSKID
		gl.FCBRANCH = glRef.FCBRANCH
		gl.FCCORP = glRef.FCCORP
		gl.FCSECT = glRef.FCSECT
		gl.FCDEPT = glRef.FCDEPT
		gl.FCGLHEAD = glhead.FCSKID
		gl.FCSEQ = fmt.Sprintf("%04d", (seqGl + 1))
		gl.FDDATE = glRef.FDDATE
		gl.FCJOB = glRef.FCJOB
		gl.FCPROJ = glRef.FCPROJ
		gl.FCCREATEBY = glRef.FCCREATEBY
		gl.FCCORRECTB = glRef.FCCREATEBY

		if runn == 0 {
			gl.FNAMT = fnSumPrice * 0.07
		} else {
			gl.FNAMT = 0 - (fnSumPrice + (fnSumPrice * 0.07))
		}

		if err := tx.Create(&gl).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		runn++
	}
	////////////////////////////////

	// fmt.Println("GLHEAD: %s", glhead.FCSKID)
	// UPDATE GLREF
	var book models.Booking
	if err := tx.First(&book, &models.Booking{FCCODE: "003", FCREFTYPE: "BI"}).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	creditDate := time.Now().AddDate(0, 0, orderH.FNCREDTERM) //time.Now() + orderH.FNCREDTERM
	glRef.FCBOOK = book.FCSKID
	glRef.FCGLHEAD = glhead.FCSKID
	glRef.FCRFTYPE = "B"
	glRef.FCREFTYPE = "BI"
	glRef.FCSTEP = "1"
	glRef.FDDUEDATE = creditDate
	glRef.FNVATAMT = vatAmt
	glRef.FNAMTKE = fnAmt
	glRef.FNVATAMTKE = fnVatAmt
	glRef.FNCREDTERM = orderH.FNCREDTERM
	glRef.FDRECEDATE, _ = time.Parse("20060102", fmt.Sprintf("%d%s", onMonth, "01"))
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
		if err := tx.First(&refProd, &models.Refprod{FCGLREF: glRef.FCSKID, FCPROD: i.FCPROD}).Error; err != nil {
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

		backQty := (i.FNBACKQTY - refProd.FNQTY)
		orderIStatus := "1"
		if backQty == 0 {
			orderIStatus = "P"
		}

		// UPDATE ORDERI
		if err := tx.Model(&models.Orderi{FCSKID: i.FCSKID}).Updates(&models.Orderi{
			FCSTEP:    orderIStatus,
			FNBACKQTY: backQty,
			FNPRICE:   refProd.FNPRICE,
		}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// fmt.Println("orderIStatus: ", orderIStatus)
		// sumCtn += orderIQty

		// // UPDATE GLREF HISTORY
		// var orderH models.Orderh
		// if err := tx.Select("FCREFNO,FNCREDTERM").First(&orderH, &models.Orderh{FCSKID: listOrderI[0].FCORDERH}).Error; err != nil {
		// 	tx.Rollback()
		// 	r.Message = err.Error()
		// 	return c.Status(fiber.StatusNotFound).JSON(&r)
		// }

		// var prod models.Product
		// if err := tx.Select("FCSKID,FCACCSCRED").First(&prod, &models.Product{FCSKID: i.FCPROD}).Error; err != nil {
		// 	tx.Rollback()
		// 	r.Message = err.Error()
		// 	return c.Status(fiber.StatusNotFound).JSON(&r)
		// }

		// CREATE NOTCUT
		var noteCut models.Notecut
		noteCut.FCBRANCH = glRef.FCBRANCH
		noteCut.FCCHILDH = listOrderI[0].FCORDERH
		noteCut.FCCHILDI = i.FCSKID
		noteCut.FCCORP = glRef.FCCORP
		noteCut.FCMASTERH = glRef.FCSKID
		noteCut.FCMASTERI = refProd.FCSKID
		noteCut.FNQTY = i.FNRECEIVEQTY
		noteCut.FCCREATEBY = refProd.FCCREATEBY
		noteCut.FCCORRECTB = glRef.FCCORRECTB
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
	// tx.Rollback()
	// store.Rollback()
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlrefHistoryController(c *fiber.Ctx) error {
	var r models.Response
	var gl []models.GlrefHistory
	if c.Query("fcskid") != "" {
		if err := configs.Store.Order("created_at").Where("fcsk_id", c.Query("fcskid")).Find(&gl).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&gl)
		}
	}

	if c.Query("fctype") != "" {
		if err := configs.Store.Limit(150).Order("created_at").Where("fctype", c.Query("fctype")).Find(&gl).Error; err != nil {
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
