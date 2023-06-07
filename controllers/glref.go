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
	glref.FCPROJ = "H2ZFfQ02"
	glref.FCJOB = frm.Job
	glref.FCCOOR = frm.Coor
	glref.FCCORP = frm.Corp
	glref.FCDEPT = frm.Department
	glref.FCSECT = sect.FCSKID
	glref.FCBOOK = book.FCSKID
	glref.FCCORRECTB = fmt.Sprintf("%s", uid)
	glref.FMMEMDATA = strings.ToUpper(frm.InvoiceNo)
	glref.FCTOWHOUSE = frm.Whs

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
		refProd.FNUMQTY = float64(i.PACK)
		refProd.FNQTY = i.Qty
		// refProd.FCUM = v.FCUM
		// refProd.FCUMSTD = v.FCUM
		// refProd.FCSTUM = v.FCUM
		// refProd.FCSTUMSTD = v.FCUM
		refProd.FCUM = i.UMID
		refProd.FCUMSTD = i.UMID
		refProd.FCSTUM = i.UMID
		refProd.FCSTUMSTD = i.UMID
		refProd.FNSTUMQTY = float64(i.PACK)
		refProd.FTDATETIME = time.Now()
		refProd.FTLASTEDIT = time.Now()
		refProd.FTLASTUPD = time.Now()
		seq++
	}

	tx.Commit()
	// End
	frm.InvoiceNo = strings.ToUpper(frm.InvoiceNo)
	r.Message = fmt.Sprintf("%s <> %s", fccode, uid)
	r.Data = &glref
	return c.Status(fiber.StatusCreated).JSON(&r)
}
