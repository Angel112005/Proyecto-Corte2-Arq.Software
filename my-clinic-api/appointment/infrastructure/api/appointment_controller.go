package appointment

import (
	"fmt"
	appAppointment "my-clinic-api/appointment/application"
	appointment "my-clinic-api/appointment/application"
	domainAppointment "my-clinic-api/appointment/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	createAppointment *appAppointment.CreateAppointment
	listAppointments  *appAppointment.ListAppointments
	updateAppointment *appAppointment.UpdateAppointment
	deleteAppointment *appAppointment.DeleteAppointment
	findByID *appointment.FindAppointmentByID

}

func NewController(
	ca *appAppointment.CreateAppointment,
	la *appAppointment.ListAppointments,
	ua *appAppointment.UpdateAppointment,
	da *appAppointment.DeleteAppointment,
	findByID *appAppointment.FindAppointmentByID,
) *Controller {
	return &Controller{
		createAppointment: ca,
		listAppointments:  la,
		updateAppointment: ua,
		deleteAppointment: da,
		findByID: findByID,
	}
}

// func (c *Controller) Create(ctx *gin.Context) {
//     var a domainAppointment.Appointment
//     if err := ctx.ShouldBindJSON(&a); err != nil {
//         fmt.Println("Error en ShouldBindJSON:", err) // Agregar este log
//         ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
//     fmt.Println("Datos recibidos:", a) // Log de los datos recibidos
//     if err := c.createAppointment.Execute(a); err != nil {
//         fmt.Println("Error en createAppointment.Execute:", err) // Log del error en la capa de aplicación
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
//         return
//     }
//     ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully"})
// }

// 📌 Crear una nueva cita y enviarla a RabbitMQ
func (c *Controller) Create(ctx *gin.Context) {
	var a domainAppointment.Appointment
	if err := ctx.ShouldBindJSON(&a); err != nil {
		fmt.Println("❌ Error en ShouldBindJSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("📌 Datos recibidos:", a)

	// 1️⃣ Guardar en MySQL y enviar evento a RabbitMQ
	if err := c.createAppointment.Execute(a); err != nil {
		fmt.Println("❌ Error en createAppointment.Execute:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "✅ Appointment created and sent to RabbitMQ"})
}

func (c *Controller) List(ctx *gin.Context) {
	appointments, err := c.listAppointments.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}

func (c *Controller) Update(ctx *gin.Context) {
	var a domainAppointment.Appointment
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	a.ID = parsedID

	if err := c.updateAppointment.Execute(a); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}

func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.deleteAppointment.Execute(parsedID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}

func (c *Controller) GetByID(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    appointment, err := c.findByID.Execute(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointment"})
        return
    }

    if appointment == nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
        return
    }

    ctx.JSON(http.StatusOK, appointment)
}