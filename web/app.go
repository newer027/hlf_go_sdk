package web
import (
	"github.com/gorilla/mux"
	"github.com/chainHero/heroes-service/web/controllers"
	"net/http"
	"log"
)

func Serve(app *controllers.Application) {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("POST").Path("/orders/").HandlerFunc(app.InitOrderHandler)
	sub.Methods("POST").Path("/states/").HandlerFunc(app.ChangeStateOrderHandler)
	sub.Methods("POST").Path("/positions/").HandlerFunc(app.UpdatePositionOrderHandler)
	sub.Methods("GET").Path("/delete_order/{orderId}/").HandlerFunc(app.DeleteHandler)
	sub.Methods("GET").Path("/orders/{orderId}/").HandlerFunc(app.ReadOrderHandler)
	sub.Methods("GET").Path("/order_detail/{orderId}/").HandlerFunc(app.QueryOrderDetailHandler)
	// sub.Methods("GET").Path("/positions").HandlerFunc(app.ReadOrderPositionHandler)
	// sub.Methods("POST").Path("/orders_range").HandlerFunc(app.GetOrdersByRangeHandler)
	sub.Methods("GET").Path("/orders_history/{orderId}/").HandlerFunc(app.GetHistoryForOrderHandler)
	// sub.Methods("GET").Path("/orders_broker/{brokerId}/").HandlerFunc(app.QueryOrdersByBrokerHandler)
	// sub.Methods("GET").Path("/orders_query").HandlerFunc(app.QueryOrdersHandler)
	sub.Methods("GET").Path("/orders_page/{brokerId}/{page}/").HandlerFunc(app.QueryOrdersWithPaginationHandler)
	// InitStringHashHandler
	// "fileId", "orderId", "dataUrl", "shaResult", "comment"
	sub.Methods("POST").Path("/string_hash/").HandlerFunc(app.InitStringHashHandler)
	sub.Methods("POST", "OPTIONS").Path("/file_hash/").HandlerFunc(app.InitFileHashHandler)

	sub.Methods("POST", "OPTIONS").Path("/users/").HandlerFunc(app.InitUserHandler)
	sub.Methods("POST", "OPTIONS").Path("/users/update/").HandlerFunc(app.UpdateUserHandler)
	sub.Methods("GET").Path("/users/{page}/").HandlerFunc(app.QueryUsersListHandler)
	sub.Methods("GET").Path("/user_detail/{userId}/").HandlerFunc(app.QueryUserDetailHandler)
	sub.Methods("GET").Path("/users/delete/{userId}/").HandlerFunc(app.DeleteUserHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}