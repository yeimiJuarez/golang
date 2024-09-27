//prorgama que lee un archivo llamado lista.txt. por cada linea encontrada, lee un archivo csv y posteriormente transforma la salida a un archivo xml
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abre el archivo lista.txt
	file, err := os.Open("lista.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo lista.txt:", err)
		return
	}
	defer file.Close()

	// Crea el scanner para lista.txt
	scanner := bufio.NewScanner(file)

	// Recorre las líneas de lista.txt
	for scanner.Scan() {
		// Obtiene el nombre del archivo CSV
		archivoInc := scanner.Text()
		fileName := archivoInc + ".csv"
		fmt.Println("Abriendo archivo:", fileName)

		// Abre el archivo CSV
		filecsv, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error al abrir el archivo", fileName, ":", err)
			continue // Continúa con la siguiente línea en lista.txt
		}
		defer filecsv.Close()

		// Crea el lector CSV
		reader := csv.NewReader(filecsv)

		// Lee todas las líneas del archivo CSV
		lines, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error al leer el archivo", fileName, ":", err)
			continue
		}

		// Verifica que el archivo CSV no esté vacío
		if len(lines) < 1 {
			fmt.Println("Archivo CSV vacío:", fileName)
			continue
		}

		// La primera línea contiene los encabezados
		headers := lines[0]

		// Crea el archivo XML
		xmlFile, err := os.Create(archivoInc + ".xml")
		if err != nil {
			fmt.Println("Error al crear el archivo XML:", err)
			continue
		}
		defer xmlFile.Close()

		// Recorre las líneas del archivo CSV, omitiendo la primera línea
		for _, line := range lines[1:] {
			// Convierte la línea a la estructura XML
			xmlRow := transformToXML(headers, line)
			
			// Escribe la estructura XML en el archivo XML
			_, err = xmlFile.WriteString(xmlRow)
			if err != nil {
				fmt.Println("Error al escribir en el archivo XML:", err)
				break
			}
		}
	}

	// Maneja posibles errores del scanner de lista.txt
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo lista.txt:", err)
	}
}

// Función para transformar una línea CSV a una estructura XML
func transformToXML(headers, values []string) string {
	var sb strings.Builder
	sb.WriteString("<row>")

	for i := 0; i < len(headers); i++ {
		sb.WriteString(fmt.Sprintf(`<field code="%s" value="%s"/>`, headers[i], values[i]))
	}

	sb.WriteString("</row>")
	return sb.String()
}
