package appointment

import (
    "encoding/json"
    "log"

    amqp "github.com/rabbitmq/amqp091-go"
)

// Estructura del mensaje que se enviará a RabbitMQ
type CitaEvento struct {
    ID              int    `json:"id"`
    PatientName     string `json:"patient_name"`
    DoctorID        int    `json:"doctor_id"`
    AppointmentDate string `json:"appointment_date"`
    Status          string `json:"status"`
}

// Publicar un mensaje en RabbitMQ
func PublicarCitaEvento(evento CitaEvento) {
    conn, err := amqp.Dial("amqp://agmc:112005@54.157.231.159:5672/")
    if err != nil {
        log.Fatalf("❌ Error al conectar con RabbitMQ: %s", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("❌ Error al abrir el canal: %s", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "citas.nuevas", // Nombre de la cola
        true,           // Durable
        false,          // Auto-delete
        false,          // Exclusiva
        false,          // No-wait
        nil,            // Argumentos
    )
    if err != nil {
        log.Fatalf("❌ Error al declarar la cola: %s", err)
    }

    cuerpo, _ := json.Marshal(evento)
    err = ch.Publish(
        "",     // Exchange
        q.Name, // Routing Key
        false,  // Mandatory
        false,  // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        cuerpo,
        })
    if err != nil {
        log.Fatalf("❌ Error al publicar mensaje: %s", err)
    }

    log.Printf("📢 Evento enviado a RabbitMQ: %s", cuerpo)
}
