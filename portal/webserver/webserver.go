package webserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/pinezapple/LibraryProject20201/portal/controller/authen"
	"github.com/pinezapple/LibraryProject20201/portal/controller/doc"
	"github.com/pinezapple/LibraryProject20201/portal/controller/user"
	"github.com/pinezapple/LibraryProject20201/portal/core"
	dao "github.com/pinezapple/LibraryProject20201/portal/dao/database"
	"github.com/pinezapple/LibraryProject20201/skeleton/logger"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"
)

// WebServer booting web server by configuration
func WebServer(ctx context.Context) (fn model.Daemon, err error) {
	// get configs
	mainConf, httpServerConf := core.GetConfig(), core.GetHTTPServerConf()
	lg := core.GetLogger()

	// create admin by default
	conf := core.GetConfig()
	dao.GetUserDAO().Create(context.Background(), core.GetDB(), conf.WebServer.Secure.SipHashSum0, conf.WebServer.Secure.SipHashSum1, "root", "20wPk29bnP2kc93nb92bzEobm", "", "", "", "", "")
	/*
		// try to load TLS Config
		var tlsConf *tls.Config
		tlsConf, err = xhttp.GenServerTLSConfig(httpServerConf)
		if err != nil {
			return
		}
	*/
	// Initialize router and http server
	e, server := echo.New(), &http.Server{
		Addr: fmt.Sprintf(":%d", httpServerConf.Port),
	}

	// Disable echo logging
	e.Logger.SetOutput(ioutil.Discard)

	// Recover middleware recovers from panics anywhere in the chain, prints stack trace and handles the control to the centralized HTTPErrorHandler.
	// Default stack size for trace is 4KB. For more example, please refer to https://echo.labstack.com/
	e.Use(mw.Recover())

	// Remove trailing slash middleware removes a trailing slash from the request URI.
	e.Pre(mw.RemoveTrailingSlash())

	// Set BodyLimit Middleware. It will panic if fail. For more example, please refer to https://echo.labstack.com/
	e.Use(mw.BodyLimit(mainConf.WebServer.BodyLimit))

	// Secure middleware provides protection against cross-site scripting (XSS) attack, content type sniffing, clickjacking, insecure connection and other code injection attacks.
	// For more example, please refer to https://echo.labstack.com/
	e.Use(mw.Secure())
	e.Use(mw.CORS())
	// Restricted group of URIs for
	/*
		r := e.Group("/r")
		r.Use(mw.JWTWithConfig(mw.JWTConfig{
			Claims:      &model.Claim{},
			ContextKey:  mainConf.WebServer.Secure.JWT.ContextKey,
			SigningKey:  []byte(mainConf.WebServer.Secure.JWT.SecretKey),
			TokenLookup: "header:authorization",
		}))
	*/

	// init router
	//initRouter(e, r)

	e.POST("/p/login", authen.Login)
	initUserRouter(e)
	initDocRouter(e)

	//	core.Logger().Infof("HTTP Server is starting on :%d", httpServerConf.Port)
	logger.LogInfo(lg, "HTTP Server is starting on "+strconv.Itoa(httpServerConf.Port))
	// Start server
	go func() {
		if err := e.StartServer(server); err != nil {
			logger.LogErr(lg, err)
			//core.Logger().Error(err)
		}
	}()

	fn = func() {
		<-ctx.Done()

		// try to shutdown server
		if err := e.Shutdown(context.Background()); err != nil {
			logger.LogErr(lg, err)
			//	core.Logger().Error(err)
		} else {
			//	core.Logger().Warn("gracefully shutdown webserver")
			logger.LogInfo(lg, "gracefully shutdown webserver")
		}
	}
	return
}

func initDocRouter(e *echo.Echo) {
	docG := e.Group("/document")
	{
		docG.POST("/create", doc.SaveDocumentByBatch)
		docG.POST("/update", doc.UpdateDocument)
	}

	docVerG := e.Group("/documentversion")
	{
		docVerG.POST("/create", doc.CreateDocumentVersion)
		docVerG.POST("/update", doc.UpdateDocumentVersion)
		docVerG.POST("/addbarcode", doc.AddBarcodeUpdateDocumentVersion)
	}

	d := e.Group("/barcode")
	d.POST("/avail", doc.SelectAllAvailableBarcode)
	d.POST("/selling", doc.SelectAllSellingBarcode)
	d.POST("/damaged", doc.SelectAllDamagedBarcode)
	d.POST("/detail", doc.SelectBarcodeByID)
	d.POST("/update", doc.UpdateBarcodeStatus)
	d.POST("/delete", doc.DeleteBarcode)

	f := e.Group("/borrow")
	f.POST("/all", doc.SelectAllBorrowForm)
	f.POST("/unreturned", doc.SelectAllNotReturnedBorrowForm)
	f.POST("/detail", doc.SelectBorrowFormByID)
	f.POST("/save", doc.CreateBorrowForm)
	f.POST("/create", doc.CreateBorrowForm)
	f.POST("/update", doc.UpdateBorrowForm)

	g := e.Group("/payment")
	g.POST("/all", doc.SelectAllPayment)
	g.POST("/detail", doc.SelectPaymentByID)

	h := e.Group("/salebill")
	h.POST("/create", doc.CreateSaleBill)
	h.POST("/all", doc.SelectAllSaleBill)
	h.POST("/detail", doc.SelectSaleBillByID)

	k := e.Group("/blacklist")
	k.POST("/all", doc.SelectAllBlackList)
	k.POST("/detail", doc.SelectBlackListByID)

	l := e.Group("/doc")
	l.POST("/all", doc.SelectAllDoc)
	l.POST("/detail", doc.SelectDocByID)

	e.POST("/docver/detail", doc.SelectDocVerByID)

	/*
		d.POST("/save", doc.SaveDoc)
		d.POST("/delete", doc.DelDoc)
		d.POST("/update", doc.UpdateDoc)
		d.POST("/alldoc", doc.SelectAllDoc)
		d.POST("/alldoc0", doc.SelectAllDoc0)
		d.POST("/onedoc", doc.SelectDocByID)

		d.POST("/saveForm", doc.SaveForm)
		d.POST("/allform", doc.SelectAllForm)
		d.POST("/oneform", doc.SelectFormByID)

		d.POST("/updateStatus", doc.UpdateStatus)
	*/
}

func initUserRouter(e *echo.Echo) {
	u := e.Group("/user")
	u.POST("/save", user.SaveUser)
	u.POST("/delete", user.DeleteUser)
	u.POST("/update", user.UpdateUser)
	u.POST("/alluser", user.SelectAllUser)
	u.POST("/oneuser", user.SelectUserByID)
}

func initRouter(e *echo.Echo, r *echo.Group) {

	/*
		initUserRouter(e)
		initGroupRouter(e)
	*/

	e.Static("/www", "www")
	// index
	e.GET("/", func(c echo.Context) error { return c.File("www/dist/index.html") })
	// login
	e.POST("/p/login", authen.Login)
}
