package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/Velocidex/ordereddict"
)

var fallos int

func check(descripcion string, ok bool) {
	if ok {
		fmt.Println("  [OK]   ", descripcion)
	} else {
		fmt.Println("  [FALLO]", descripcion)
		fallos++
	}
}

func probarStack() {
	fmt.Println("=== Stack ===")

	// crear un stack; recién creado debe estar vacío
	s := NewStack()
	check("stack recién creado está vacío", s.isEmpty())

	// pop en un stack vacío devuelve false
	_, ok := s.pop()
	check("pop en stack vacío devuelve false", !ok)

	// push de 3 elementos
	s.push(10)
	s.push(20)
	s.push(30)
	check("después de push, isEmpty devuelve false", !s.isEmpty())

	// peek devuelve el último elemento pero no lo saca
	v, ok := s.peek()
	check("peek devuelve 30", ok && v == 30)
	v, ok = s.peek()
	check("peek no saca el elemento (sigue siendo 30)", ok && v == 30)

	// al sacar los elementos salen en orden inverso
	esperado := []int{30, 20, 10}
	for _, e := range esperado {
		v, ok := s.pop()
		check(fmt.Sprintf("pop devuelve %d", e), ok && v == e)
	}
	check("después de sacar todo, el stack está vacío", s.isEmpty())
}

func probarQueue() {
	fmt.Println("=== Queue ===")

	// crear un queue; recién creado debe estar vacío
	q := NewQueue()
	check("queue recién creado está vacío", q.isEmpty() && q.size() == 0)

	// dequeue en un queue vacío devuelve false
	_, ok := q.dequeue()
	check("dequeue en queue vacío devuelve false", !ok)

	// enqueue de 3 elementos
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)

	// front, back y size devuelven el valor correcto
	v, ok := q.front()
	check("front devuelve 10", ok && v == 10)
	v, ok = q.back()
	check("back devuelve 30", ok && v == 30)
	check("size devuelve 3", q.size() == 3)

	// al hacer dequeue salen en el mismo orden en que entraron
	esperado := []int{10, 20, 30}
	for _, e := range esperado {
		v, ok := q.dequeue()
		check(fmt.Sprintf("dequeue devuelve %d", e), ok && v == e)
	}
	check("después de sacar todo, el queue está vacío", q.isEmpty())
}

func probarDict() {
	fmt.Println("=== Dict (ordereddict) ===")

	// crear un diccionario nuevo; recién creado debe estar vacío
	d := ordereddict.NewDict()
	check("dict recién creado está vacío", d.Len() == 0)

	// get de una clave que no existe devuelve false
	_, ok := d.Get("uno")
	check("get de una clave que no existe devuelve false", !ok)

	// agregar nuevos key-value pairs con Set
	d.Set("uno", 1)
	d.Set("dos", 2)
	d.Set("tres", 3)
	check("después de agregar 3 pares, Len devuelve 3", d.Len() == 3)
	check("las claves quedan en orden de inserción [uno dos tres]",
		slices.Equal(d.Keys(), []string{"uno", "dos", "tres"}))

	// lookup: Get devuelve el valor guardado para la clave
	v, ok := d.Get("dos")
	check("get(\"dos\") devuelve 2", ok && v == 2)

	// update: Update cambia el valor sin mover la clave de lugar
	d.Update("dos", 22)
	v, ok = d.Get("dos")
	check("después de update, get(\"dos\") devuelve 22", ok && v == 22)
	check("update no cambia el tamaño (Len sigue en 3)", d.Len() == 3)
	check("update no cambia el orden de las claves",
		slices.Equal(d.Keys(), []string{"uno", "dos", "tres"}))

	// delete: Delete elimina la clave indicada
	d.Delete("uno")
	_, ok = d.Get("uno")
	check("después de delete, get(\"uno\") devuelve false", !ok)
	check("después de delete, Len devuelve 2", d.Len() == 2)
	check("las demás claves conservan su orden [dos tres]",
		slices.Equal(d.Keys(), []string{"dos", "tres"}))

	// iterar sobre los elementos del dict (en orden)
	var claves []string
	var valores []any
	for _, item := range d.Items() {
		fmt.Printf("           %s = %v\n", item.Key, item.Value)
		claves = append(claves, item.Key)
		valores = append(valores, item.Value)
	}
	check("la iteración recorre las claves en orden [dos tres]",
		slices.Equal(claves, []string{"dos", "tres"}))
	check("la iteración devuelve los valores correctos [22 3]",
		slices.Equal(valores, []any{22, 3}))
}

func main() {
	probarStack()
	fmt.Println()
	probarQueue()
	fmt.Println()
	probarDict()
	fmt.Println()

	if fallos > 0 {
		fmt.Printf("%d comprobación(es) fallaron\n", fallos)
		os.Exit(1)
	}
	fmt.Println("Todas las comprobaciones pasaron")
}
