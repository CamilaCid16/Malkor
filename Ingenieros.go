package main

import (
	"bufio"
	"context"
	"fmt"
	"log"


	pb "github.com/yojeje/lab6"
	"google.golang.org/grpc"
)

var (
	nombre      string //Nombre (registro)
	relojx      int64    //Dimension X del reloj de vector
	relojy      int64    //Dimension Y del reloj de vector
	relojz      int64    //Dimension Z del reloj de vector
	lastfulcrum string //ip del ultimo fulcrum consultado para este planeta
)


func Consultar(m pb.IngenierosClient) {
	var opcion int;
	var comando string;
	fmt.Println("1) Agregar Base\n2) Actualizar Valor\n3) Renombrar Base\n4) Borrar Base\nSeleccione un comando:")
	fmt.Scanln(&opcion)

	if (opcion == 1) {
		comando = "AgregarBase"
	} else if (opcion == 2) {
		comando = "ActualizarValor"
	} else if (opcion == 3) {
		comando = "RenombrarBase"
	} else if (opcion == 4) {
		comando = "BorrarBase"
	}
	response, err := m.EnviarBroker(context.Background(), &pb.Comando{Tipo: comando})

	if (err != nil) {
		log.Fatalf("Error %v", bufio.ErrBadReadCount)
	}


	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial(response.Dir, grpc.WithInsecure());
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	s := pb.NewIngenierosClient(conn)
	fmt.Println("Enviando comando al servidor" + response.Dir)
	responseS, err := s.EnviarServidor(context.Background(), &pb.Comando{Tipo: comando})

	relojx += responseS.X
	relojy += responseS.Y
	relojz += responseS.Z
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