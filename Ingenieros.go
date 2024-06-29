package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	pb "github.com/yojeje/lab6"
	"google.golang.org/grpc"
)


func Consultar(m pb.IngenierosClient) {
	var opcion string
	var comando string
	scanner := bufio.NewScanner(os.Stdin)

	// Leer la opción del usuario
	fmt.Println("1) Agregar Base\n2) Actualizar Valor\n3) Renombrar Base\n4) Borrar Base\nSeleccione un comando:")
	for scanner.Scan() {
		opcion = scanner.Text()
		if opcion != "" {
			break
		}
		fmt.Println("Input cannot be empty. Please enter a command:")
	}
	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}

	// Asignar el comando correspondiente según la opción seleccionada
	switch opcion {
	case "1":
		comando = "AgregarBase"
	case "2":
		comando = "ActualizarValor"
	case "3":
		comando = "RenombrarBase"
	case "4":
		comando = "BorrarBase"
	default:
		fmt.Println("Opción inválida:", opcion)
		return
	}

	// Dividir el comando si es necesario (ejemplo con strings.Split)
	comand := strings.Split(comando, " ")
	fmt.Println("Comando:", comand[0])

	// Aquí debes llamar a la función para enviar el comando al broker
	// Asegúrate de ajustar `m.EnviarBroker` y `pb.Comando` según tus definiciones reales
	response, err := m.EnviarBroker(context.Background(), &pb.Comando{Tipo: comando})
	if err != nil {
		log.Fatalf("Error al enviar comando: %v", err)
	}
	
	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial(response.Dir, grpc.WithInsecure());
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	s := pb.NewIngenierosClient(conn)
	fmt.Println("Enviando comando al servidor" + response.Dir)
	responseS, err := s.EnviarServidor(context.Background(), &pb.Comando{Tipo: comando})

	if err != nil {
		log.Println("ERROR")
	}
	
	fmt.Println(responseS.X, responseS.Y, responseS.Z)
}

func main() {
	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial("dist098:50051", grpc.WithInsecure());
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	m := pb.NewIngenierosClient(conn)
	Consultar(m)
	defer conn.Close()
}