// Aula 2 - Hello World

/*
package main

import (
	"fmt"
)
func main() {
	fmt.Printf("Hello, World!\n")
} /*/

// Aula 3, 4 e 5- Variáveis, Valores e Tipos de Dados

// Variáveis são usadas para armazenar valores que podem ser alterados durante a execução do programa.
// Em Go, você pode declarar variáveis usando a palavra-chave 'var' ou a atribuição curta ':='.
// A declaração 'var' é usada para declarar uma variável com um tipo específico, enquanto ':=' é usado para inferir o tipo automaticamente.
// Você também pode declarar múltiplas variáveis em uma única linha, separando-as por vírgulas.
// Além disso, você pode declarar constantes usando a palavra-chave 'const', que são valores que não podem ser alterados após a declaração.

/* func main() {

	var nome string = "John"
	fmt.Println("Nome:", nome)

	idade := 30
	fmt.Println("Idade:", idade)

	var cidade, país = "São Caetano do Sul", "Brasil"
	fmt.Println("Cidade:", cidade, "\nPaís:", país)

	const pi = 3.14
	fmt.Println("Pi:", pi)
} /*/

// Tipo Float64
/*
var x = 10.5

func main() {
	fmt.Printf("%v, %T", x, x)
} /*/

// Tipo Int
/*
var x = 10

func main() {
	fmt.Printf("%v, %T", x, x)
} /*/

// Tipo String
/*
var x = "3CON - Go"

func main() {
	fmt.Printf("%v, %T", x, x)
} /*/

// Tipo Boolean
/*
var x = true
var y = false

func main() {
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
}/*/

// Aula 6 - Valores Zero

// Valores zero são os valores padrão atribuídos a variáveis quando elas são declaradas, mas não inicializadas.
// Para tipos numéricos, o valor zero é 0; para strings, é uma string vazia ""; e para booleanos, é false.
// Esses valores são úteis para evitar erros quando você tenta usar uma variável antes de atribuir um valor a ela.
/*
var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
} /*/

// Aula 7 - "fmt"
// O pacote "fmt" é usado para formatar e imprimir dados no console.
// Ele fornece funções como Println, Printf e Sprintf para exibir informações de forma formatada.
// Acento grave (`) é usado para criar strings multilinha ou para evitar a necessidade de escapar de caracteres especiais.

/*
func main() {
	x := `"O acento grave é usado para \n criar

	\n strings multilinha ou para evitar %v a necessidade

	de escapar de caracteres especiais."`

	fmt.Println(x)
} /*/

// Sprint é uma função do pacote "fmt" que formata os dados como uma string.
// Ela é útil quando você deseja criar uma string formatada sem imprimi-la diretamente no console
/*
func main() {
	x := "3Con"
	y := " IT & Digital"

	z := fmt.Sprint(x, y)

	fmt.Println(z)
} /*/

/*
	Grupo #1 - Print --> Imprime no console
	print: Imprime sem formatação, semelhante a println, mas sem quebra de linha.
	println: Imprime uma linha com quebra de linha no final.
	printf: Imprime formatando a string com especificadores de formato.

	Grupo #2 - Sprint --> Retorna uma string formatada
	sprint: Retorna uma string formatada sem imprimir no console.
	sprintln: Retorna uma string formatada com quebra de linha no final.
	sprintf: Retorna uma string formatada com especificadores de formato.

	Grupo #3 - Fprint --> Imprime em um Writer
	fprint: Imprime em um Writer sem formatação.
	fprintln: Imprime em um Writer com quebra de linha no final.
	fprintf: Imprime em um Writer formatando a string com especificadores de formato.

/
*/

// Aula 10 - Exercicios de Fixação
// Exercicio 1: Declare e imprima variáveis de diferentes tipos
// Exemplo de declaração e impressão de variáveis de diferentes tipos
/*
import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
/*/

// Exercicio 2: Declare variáveis com valores zero e imprima
// Exemplo de declaração de variáveis com valores zero e impressão
/*
package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {

	fmt.Println(x, y, z)
}
/*/

// Exercicio 3: Use fmt.Sprintf para formatar uma string com variáveis
// Aqui, usamos fmt.Sprintf para formatar uma string com variáveis de diferentes tipos.
/*
package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main () {
s := fmt.Sprintf("%v, %v, %v", x, y, z)
	fmt.Println(s)
}
/*/

// Exercicio 4: Declare uma variável personalizada e crie um tipo baseado nela
// Neste exemplo, criamos um tipo personalizado chamado "thiago" e declaramos uma variável desse tipo.
/*
package main

import (
	"fmt"
)
type thiago int

var x thiago

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Printf("%v", x)
}
/*/

// Exercicio 5: Atribuição, Package-Level-Scoped Variaveis.
// Neste exemplo, declaramos duas variáveis do tipo personalizado "thiago" e atribuímos valores a elas.
// Var x e y são variáveis de nível de pacote, o que significa que elas podem ser acessadas em todo o pacote.
// Dentro da função main, atribuímos o valor de x a y e imprimimos ambos os valores e seus tipos.

/*
package main

import (
	"fmt"
)

type thiago int

var x thiago
var y thiago

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = x
	fmt.Printf("%v, %T\n", y, y)
}

/*/

//Aula 22 - Tipo Booleano
// O tipo booleano em Go é usado para representar valores lógicos, que podem ser true (verdadeiro) ou false (falso).
// Ele é frequentemente usado em condições de controle de fluxo, como if, for e switch.
// Operadores relacionais, como ==, !=, <, >, <= e >=, retornam valores booleanos.

/*
package main

import (
	"fmt"
)

var x bool
var a int
var y int
var z bool

func main() {
	fmt.Println(x)   // Imprime o valor inicial de x, que é false (valor zero para booleano)
	x = true         // Atribui o valor true à variável x
	fmt.Println(x)   // Imprime o novo valor de x, que agora é true
	x = (10 < 100)   // Atribui o resultado da expressão (10 < 100) a x, que será true
	fmt.Println(x)   // Imprime o valor de x, que ainda é true
	y := (10 == 100) // Declara uma nova variável y e atribui o resultado da expressão (10 == 100), que será false
	fmt.Println(y)   // Imprime o valor de y, que é false

}
/*/

// **Aula 23 - Tipos Numericos**
// Em Go, os tipos numéricos são usados para representar valores numéricos, como inteiros e números de ponto flutuante.
// Existem vários tipos numéricos, incluindo int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32 e float64.
// uint representa números inteiros sem sinal, enquanto int representa números inteiros com sinal. (uint8, uint16, uint32 e uint64 são usados para representar inteiros sem sinal de tamanhos específicos).
// float32 e float64 são usados para representar números de ponto flutuante, com float 32 tendo precisão menor que float64.
// int é o tipo padrão para números inteiros em Go, e seu tamanho depende da plataforma (32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits).
// Cada tipo numérico tem um tamanho específico e é usado para armazenar diferentes faixas de valores. (int serve para armazenar números inteiros, int8, int16, int32 e int64 são usados para armazenar inteiros de tamanhos específicos.)

/*
package main
import (
	"fmt"
)

func main() {

	var x int = 10
	var y float64 = 3.14
	var z uint = 42

	fmt.Printf("x: %v, Tipo: %T\n", x, x)
	fmt.Printf("y: %v, Tipo: %T\n", y, y)
	fmt.Printf("z: %v, Tipo: %T\n", z, z)
}
/*/

// **Aula 23 - Overflow em Tipos Numéricos**
// Overflow ocorre quando um valor excede o limite máximo ou mínimo de um tipo numérico, resultando em um comportamento inesperado.
// Em Go, os tipos numéricos têm limites específicos, e quando um valor ultrapassa esses limites, ele "transborda" para o outro extremo do intervalo permitido.
// Por exemplo, se você tentar armazenar um valor maior que o máximo permitido para um tipo, ele retornará ao valor mínimo permitido, e vice-versa.

/*
package main
import (
	"fmt"
)

func main() {
	var i uint16
	i = 65535 // Valor máximo para uint16
	fmt.Println("Valor máximo de uint16:", i)

	i++ // Isso causará overflow, retornando ao valor mínimo
	fmt.Println("Após overflow, valor de uint16:", i) // Imprime 0, i retorna ao valor mínimo permitido (0) devido ao overflow

	i = 65536 // Tentando atribuir um valor maior que o máximo permitido
	fmt.Println("Após atribuição, valor de uint16:", i) // Imprime um erro, pois 65536 não pode ser armazenado em um uint16, resultando em um overflow e retornando ao valor mínimo (0) novamente.

}
/*/

// Aula 28 - Constantes
// Constantes são valores que não podem ser alterados após serem definidos. Elas são úteis para definir valores fixos que não devem ser modificados durante a execução do programa.
// Em Go, você pode declarar constantes usando a palavra-chave `const`. As constantes podem ser de qualquer tipo, incluindo números, strings e booleanos.
// As constantes são frequentemente usadas para definir valores que serão usados em todo o programa, como limites, taxas de juros, ou outras configurações que não devem ser alteradas.

/*
package main

import (
	"fmt"
)

const x = 10 // Constante do tipo int
var y = 10 // Variável do tipo int

var c int
var d float64

func main() {
	d = x
	fmt.Println(d)
}
	/*/

// Aula 29 - iota
// O `iota` é uma constante especial em Go que é usada para criar sequências de valores constantes.
// Ele é frequentemente usado para gerar valores incrementais automaticamente, especialmente em enumerações.
// O `iota` começa em 0 e é incrementado automaticamente a cada linha onde é usado dentro de um bloco de constantes.

/*
package main

import (
	"fmt"
)

const (
	x = iota // x será 0
	y = iota // y será 1
	z = iota // z será 2
	_ = iota // _ será 3, mas não será usado
	// Você pode pular valores usando o _ (underscore) para ignorar o valor de iota.
	// Você pode usar o iota para criar constantes com valores incrementais.
	// Exemplo: iota pode ser usado para criar constantes de status, como StatusOK = iota // StatusOK será 4
	thiago = iota // thiago será 4
)

func main() {
	fmt.Println(x, y, z, thiago) // Imprime: 0 1 2
}
/*/

// Aula 30 - Deslocamento de Bits
// O deslocamento de bits é uma operação que move os bits de um número para a esquerda ou para a direita.
// Em Go, você pode usar os operadores `<<` (deslocamento para a esquerda) e `>>` (deslocamento para a direita) para realizar essas operações.
// O deslocamento para a esquerda multiplica o número por 2 elevado ao número de posições deslocadas, enquanto o deslocamento para a direita divide o número por 2 elevado ao número de posições deslocadas.

// No exemplo abaixo, o número 24 é deslocado para a direita em 2 posições, resultando em 6 (24 / 4), e deslocado para a esquerda em 2 posições, resultando em 96 (24 * 4). Os resultados são impressos em formato binário usando `%b` no `fmt.Printf`. A operação de deslocamento é útil em várias aplicações, como manipulação de bits, criptografia e otimização de algoritmos. Com o deslocamento de bits, você pode realizar operações eficientes em números inteiros, como multiplicação e divisão por potências de 2, sem usar operações aritméticas convencionais.

/*
package main

import (
	"fmt"
)

func main() {
	y := 24

	x := y >> 2

	z := y << 2

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
}
/*/

// Aula 31 - Exercicios de Fixação
// Exercicio 1 - Mostre numeros decimais, binários e hexadecimais.

/*
package main

import (
	"fmt"
)

func main() {

	var x int = 10
	var y int = 20
	var z int = 30

	fmt.Printf("Decimal: %d, Binário: %b, Hexadecimal: %#x\n", x, x, x)
	fmt.Printf("Decimal: %d, Binário: %b, Hexadecimal: %#x\n", y, y, y)
	fmt.Printf("Decimal: %d, Binário: %b, Hexadecimal: %#x\n", z, z, z)
}
/*/

// Exercicio 2 - Utilize expressões dos seguintes operadores: ==, <=, >=, !=, > e <.

/*
package main

import (
	"fmt"
)

func main() {

	var x int = 10
	var y int = 20

	fmt.Printf("%d == %d", x, y)
	s := x == y
	fmt.Println("\n", s)
	fmt.Printf("\n%d >= %d", x, y)
	a := x >= y
	fmt.Println("\n", a)
	fmt.Printf("\n%d <= %d", x, y)
	b := x <= y
	fmt.Println("\n", b)
	fmt.Printf("\n%d > %d", x, y)
	c := x > y
	fmt.Println("\n", c)
	fmt.Printf("\n%d < %d", x, y)
	d := x < y
	fmt.Println("\n", d)
}
/*/

// Exercicio 3 - Crie constantes tipadas e não tipadas, e imprima seus valores e tipos.

/*
package main

import (
	"fmt"
)

const x = 10 // Constante não tipada
const y int = 20 // Constante tipada

func main() {

	fmt.Println(x)
	fmt.Printf("%v, %T", y, y)

}
/*/

// Exercicio 4 - Crie uma variavel valor int, e demonstre em binário, decimal e hexadecimal, e faça deslocamento de bits dessa variavel 1 para esquerda e atribua o resultado a outra variavel. E depois demonstre o resultado em binário, decimal e hexadecimal.

/*
package main

import (
	"fmt"
)

var x int
var y int

func main() {

	x = 20
	fmt.Printf("Decimal: %d, Binário: %b, Hexadecimal: %#x\n", x, x, x)
	y = x << 1 // Deslocamento de bits para a esquerda
	fmt.Printf("Deslocamento para a esquerda: Decimal: %d, Binário: %b, Hexadecimal: %#x\n", y, y, y)

}
/*/

// Exercicio 5 - Criar uma variavel do tipo String, utilizando raw string literal, e imprima o valor dela no console.

/*
package main

import (
	"fmt"
)

var thiago = `3CON - Go
Thiago %2632¨@¨¨@@&&@@**# \n

Aula 5 - Variáveis, Valores e Tipos de Dados}}}}`

func main() {
	fmt.Println(thiago)
}
/*/

// Exercicio 6 - Crie 4 constantes cujo os valores sejam os proximos 4 anos, usando iota, e imprima os valores delas no console.

/*
package main

import (
	"fmt"
)

const (
	anoAtual =  2025
	proximoAno = iota + anoAtual
	segundoAno = iota + anoAtual
	terceiroAno = iota + anoAtual
	quartoAno = iota + anoAtual
)

func main() {
	fmt.Println(proximoAno, segundoAno, terceiroAno, quartoAno)
}
/*/

// Aula 38 - Fluxo de controle ( Loops: inicialização, condição e pós-ação)

/*
package main

import (
	"fmt"
)

func main() {

	// Loop for com inicialização, condição e pós-ação
	for i := 0; i < 5; i++ {
		fmt.Println("Iteração:", i)
	}

	// Loop while (usando for sem inicialização e pós-ação)
	j := 0
	for j < 5 {
		fmt.Println("Iteração while:", j)
		j++
	}

	// Loop infinito (cuidado ao usar!)
	// k := 0
	// for {
	// 	fmt.Println("Loop infinito:", k)
	// 	k++
	// 	if k >= 5 {
	// 		break // Para evitar um loop infinito, usamos break para sair do loop
	// 	}
	// }

}
/*/

// Aula 43 - Desafio Surpresa ( Faça um loop dos numero 33 a 122, e utilize format printing para demonstra-los com string/texto.)
/*
package main

import (
	"fmt"
)

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("Número: %d, Caractere: %c, Binário: %b, Decimal: %d, Hexadecimal: %#x\n", i, i, i, i, i)
	}
}
/*/

//Aula 44 - Condicionais (if, else if, else)
// Condicionais são estruturas de controle que permitem executar diferentes blocos de código com base em condições específicas.
// Em Go, você pode usar `if`, `else if` e `else` para controlar o fluxo do programa com base em expressões booleanas.

/*
package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 20

	if x < y {
		fmt.Println("x é menor que y")
	} else if x > y {
		fmt.Println("x é maior que y")
	} else {
		fmt.Println("x é igual a y")
	}

	// Exemplo de uso de operadores lógicos
	if x < 15 && y > 15 {
		fmt.Println("x é menor que 15 e y é maior que 15")
	}

	if x == 10 || y == 20 {
		fmt.Println("x é igual a 10 ou y é igual a 20")
	}
}
/*/

// Aula 45 - Switch
// O `switch` é uma estrutura de controle que permite executar diferentes blocos de código com base no valor de uma expressão.
// Ele é uma alternativa mais legível e organizada ao uso de múltiplos `if` e `else if`.

/*
package main

import (
	"fmt"
)

func main() {
	x := 10 // Declarando uma variável x com o valor 10, usando uma atribuição curta.

	switch x {
	case 5: //O "case" significa que se x for igual a 5, o bloco de código abaixo será executado.
		fmt.Println("x é igual a 5")
	case 10:
		fmt.Println("x é igual a 10")
	case 15:
		fmt.Println("x é igual a 15")
	default: //O "default" é executado se nenhum dos casos anteriores for verdadeiro.
		fmt.Println("x não é igual a nenhum dos casos anteriores")
	}

	// Exemplo de switch com múltiplos casos
	switch x {
	case 5, 10, 15:
		fmt.Println("x é 5, 10 ou 15")
	default:
		fmt.Println("x não é 5, 10 ou 15")
	}
}
/*/

// Aula 48 - Operadores Lógicos Condicionais
// Os operadores lógicos condicionais são usados para combinar ou inverter condições booleanas.
// Em Go, os principais operadores lógicos são `&&` (E lógico), `||` (OU lógico) e `!` (NÃO lógico).

/*
package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 20

	if x < y && y > 15 { // Verifica se x é menor que y e y é maior que 15
		fmt.Println("x é menor que y e y é maior que 15")
	}

	if x > 5 || y < 25 { // Verifica se x é maior que 5 ou y é menor que 25
		fmt.Println("x é maior que 5 ou y é menor que 25")
	}

	if !(x == y) { // Verifica se x não é igual a y
		fmt.Println("x não é igual a y")
	}

}
/*/

// Aula 49 - Exercícios de Fixação
// Exercicio 1 - Imprima todos os numero 1 á 10000.
/*
package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}
/*/
// Exercicio 2 - Põe na tela : O unicode  code point de todas as letras maiúsculas do alfabeto, tres vezes cada.

/*
package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ { // Letras maiúsculas de A (65) a Z (90)
		fmt.Println("\n", i) // Imprime o valor numérico de cada letra maiúscula
		for j := 0; j < 3; j++ { // Imprime três vezes cada letra
			fmt.Printf("\t %#U\n", i) // Imprime o code point Unicode da letra maiúscula
		}
	}
}
/*/

// Exercicio 3 - Criar um loop  utilizando a sintaxe: for condition {}.
/*
package main

import (
	"fmt"
)

func main() {

	nascimento := 2005
	anoAtual := 2025

	for nascimento <= anoAtual {
		fmt.Println(nascimento)
		nascimento ++
	}
}
/*/

// Exercicio 4 - Criar um loop usando a sintaxe for{}
/*
package main

import (
	"fmt"
)

func main() {

	nascimento := 2005
	anoAtual := 2025

	for {
		if nascimento > anoAtual {
			break // Não tirar o break
		}
		fmt.Println(nascimento)
		nascimento++
	}
}
/*/

// Exercicio 5 - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
/*
package main

import (
	"fmt"
)

func main () {

	for i := 0; i <= 100; i++ {
		if i % 4 == 0 {
			fmt.Println("Este número é divisivel por 4: ", i)
		}
	}
}
/*/

//Exercicio 6 - Crie um programa que demonstra o funcionamento da declaração if.

/*
package main

import (
	"fmt"
)

func main() {

	var idade int = 11

	if idade < 18 {
		fmt.Printf("O jovem não pode entrar!")
	} else {
		fmt.Printf("Pode entrar")
	}
}
/*/

// Exercicio 7 - Crie um programa que demonstra o funcionamento de else if e else.

/*
package main

import (
	"fmt"
)

func main() {

	var idade int = 21

	if idade < 18 {
		fmt.Printf("O jovem não pode ser escalado!")
	} else if idade > 20 {
		fmt.Printf("O jovem pode ser escalado e jogar")
	} else {
		fmt.Printf("O jovem pode jogar, mas não será escalado")
	}
}
/*/

//Exercicio 8 - Crie um programa que utilize a declaração switch, sem switch statement especificado.
/*
package main

import (
	"fmt"
)

func main() {

	var idade int = 14

	switch idade {
	case 14:
		fmt.Printf("O jogador é sub-14")
	case 16:
		fmt.Printf("O jogador é sub-16")
	case 18:
		fmt.Printf("O jogador é sub-18")
	case 20:
		fmt.Printf("O jogador é sub-20")
	case 21:
		fmt.Printf("O jogador é Profissional")
	default:
		fmt.Printf("Ele não é jogador!")
	}

	switch idade {
	case 14, 16, 18, 20, 21:
		fmt.Printf("\nEle é jogador!")
	default:
		fmt.Printf("\nEle não é jogador!")
	}
}
/*/

// Exercicio 9 - Crie um programa  que utilize a declaração  switch, onde o switch statement seja uma variavel do tipo string e identificador "esporteFavorito".

/*
package main

import (
	"fmt"
)

func main() {

	var esporteFavorito string = "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Printf("Ele gosta de Futebol!")
	case "Volei":
		fmt.Printf("Ele gosta de volei")
	}

}
/*/

//Exercicio 10 - Escreva resultados antes de rodar da seguinte função: fmt.Println(true && true)
/* fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

package main

import (
	"fmt"
)

func main() {

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
/*/

// Aula 59 - Agrupamentos de dados - Array

/*
package main

import (
	"fmt"
)

var x [5] int // Quantidade de indices ou posições
var y [6] int // Quantidade de indices ou posições

func main() {

	x[0] = 1 // Primeira posição do array
	x[1] = 10 // Segunda posição
	x[2] = 20 // Terceira posição
	x[3] = 30 // Quarta posição
	x[4] = 40 // Quinta posição
	fmt.Println(x[0], x[1], x[2], x[3], x[4]) // Imprime cada posição especifica e seus valores
	fmt.Println(x) // Imprime o valor do array inteiro
	fmt.Printf("%T\n", x) //Imprime o tipo do array X
	fmt.Printf("%T\n", y) //Imprime o tipo do array Y
	fmt.Println(len(x)) //Imprime o numero de indices do array.
}
/*/

// Aula 60 - Agrupamento de dados - Slice - Literal composta

/*
package main

import (
	"fmt"
)


func main() {

	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	slice2 := append(slice, 10)
	fmt.Println(slice2)

}
/*/

// Aula 61 - Agrupamento de dados - Slice: for range

/*
package main

import (
	"fmt"
)

func main() {

	slice := []string {"bola", "chuteira", "luva"}

	for indice, valor := range slice {
		fmt.Println("No indice", indice,"tem a palavra: ", valor)
	}

	slice = append(slice, "melancia")
	for indice, valor := range slice {
		fmt.Println("No indice", indice,"tem a palavra: ", valor)
	}

	for _, valor := range slice {
		fmt.Println("Uma das palavras são: ", valor)
	}
}
/*/

// Aula 62 - Agrupamentos de dados - Slice: fatiando ou deletando de uma fatia
// Exercicio surpresa
/*
package main

import (

	"fmt"
)

func main() {
	sabores := []string {"Peperoni", "Mussarela", "Abacaxi", "4 Queijos", "Margherita"}

	fatia := sabores[:]

	fmt.Println(fatia)
	fmt.Println(sabores[0], sabores[1], sabores[2], sabores[3], sabores[4])
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}
/*/

// Aula 63 - Agrupamento de dados - Slice: Anexando a uma slice
/*
package main

import (
	"fmt"
)

func main() {
	umaslice := []int {1, 2, 3, 4}
	outraslice := []int {9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)
}
/*/

//Aula 64 - Agrupamento de dados - Slice: Make

/*
package main

import (
	"fmt"
)

func main() {

	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))
}
/*/

//Aula 65 - Agrupamento de dados - Slice: Multi - Dimensional
/*
package main

import(
	"fmt"
)

func main() {
	ss := [][]int {
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
	}
	fmt.Println(ss[1][2])

}
/*/

//Aula 66 - Agrupamento de dados - Slice: Array Subjacente
/*
package main

import (
	"fmt"
)

func main() {

	primeiroSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(primeiroSlice)

	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)

	fmt.Println(segundoSlice)
	fmt.Println(primeiroSlice)
}
/*/

//Aula 69 - Exercicios de fixação
/*
package main

import (
	"fmt"
)

func main() {

	val := [5]int {1, 2, 3, 4, 5}

	for i, v := range val {
		fmt.Println(i, v)
	}

	fmt.Printf("%T, %v", val, val)
}
/*/

//Exercicio 2 - Literal composta
/*
package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)
}
/*/

//Exercicio 3 - Slicing
/*
package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2:len(slice)-1])
}
/*/

//Exercicio 4 - Demosntrando uma Slice
/*
package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
/*/

//Exercicio 5 - Slicing & Append
/*
package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[len(x)-4:]...)
	fmt.Println(x)

}
/*/

//Exercicio 6 - Len & Cap
/*
package main

import (
	"fmt"
)

func main() {
	estados := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espirito Santo", "Goiás", "Maranhão", "Mato Grosso do Sul", "Mato Grosso", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estados), cap(estados))
	fmt.Println(estados)
}
/*/

//Exercicio 7 - Slice Multi-Dimensional
/*
package main

import (
	"fmt"
)

func main() {

	info := [][]string{
		[]string{"Thiago"},
		[]string{"Tertuliano"},
		[]string{"Jogar Minecraft"},
	}
	for _, v := range info {
		fmt.Println(v[0])
		for _, dados := range v {
			fmt.Println("\t", dados)
		}
	}
}
/*/

//Exercicio 8 - Maps
/*
package main

import (
	"fmt"
)

func main() {
	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": {
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": {
			"andar de jetski", "ser milionario", "ser paulista",
		},
		"ronaldo_cristiano": {
			"melhor do mundo", "bilionario", "jogador do real madrid",
		},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t",i," - ", h)
		}
	}

	fmt.Println(mepe)
}
	/*/

//Exercicio 9 - Maps com entrada
/*
package main

import (
	"fmt"
)

func main() {
	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": {
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": {
			"andar de jetski", "ser milionario", "ser paulista",
		},
		"ronaldo_cristiano": {
			"melhor do mundo", "bilionario", "jogador do real madrid",
		},
	}

	mepe["junior_neymar"] = []string{"dribles ágils","camisa 10", "influenciador"}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t",i," - ", h)
		}
	}

	fmt.Println(mepe)
}
	/*/

//Exercicio 10 - Remover uma entrada
/*
package main

import (
	"fmt"
)

func main() {
	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": {
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": {
			"andar de jetski", "ser milionario", "ser paulista",
		},
		"ronaldo_cristiano": {
			"melhor do mundo", "bilionario", "jogador do real madrid",
		},
	}

	mepe["junior_neymar"] = []string{"dribles ágils", "camisa 10", "influenciador"}

	delete(mepe, "dasilva_guriaestranhadostrangerthings")

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
/*/

// Aula 79 - Struct
/*
package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Thiago",
		sobrenome: "Matos",
		fumante:   false,
	}

	fmt.Println(cliente1)
}
/*/

//Aula 80 - Structs Embutidos
/*
package main

import (
	"fmt"
)

type pessoa struct{
	nome string
	idade int
}

type profissional struct {
	pessoa
	titulo string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome: "Thiago",
		idade: 20,
	}

	pessoa2 := profissional{
		pessoa: pessoa {
			nome: "Thiago",
			idade: 20,
		},
			titulo: "Desenvolvedor web",
			salario: 2000,
		}

	fmt.Println(pessoa1)
	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.salario)
}
/*/

//Aula 82 - Struct Anonimos
/*
package main

import(
	"fmt"
)

func main() {
	x := struct {
		nome string
		idade int
	}{
		nome: "Mario",
		idade: 20,
	}

	fmt.Println(x)
}
/*/

// Exercicio 1 - Struct com mais de um valor
/*
package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int
	sabores []string
}

func main() {
	pessoa1 := pessoa{
		nome: "Thiago",
		idade: 20,
		sabores: []string{"Morango", "Chocolate"},
	}

	fmt.Println(pessoa1)
}
/*/

//Exercicio 2 - Struct com map
/*
package main

import "fmt"

type pessoa struct {
	nome    string
	idade   int
	sabores []string
}

func main() {
	 meumapa := make(map[string]pessoa)

	meumapa["Renata"] = pessoa{
		nome:    "Renata",
		idade:   34,
		sabores: []string{"Chocolate", "Morango"},
	}

	for _, valor := range meumapa {
		fmt.Println(valor)
	}

}
/*/

//Exercicio 3 - Struct
/*
package main

import (
	"fmt"
)

type veiculo struct {
		portas int
		cor string
	}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
	}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	carrãodotio := sedan{veiculo{4, "laranja"}, true}
	fuscavo := caminhonete{veiculo{2, "Azul"}, true}

	fmt.Println(carrãodotio)
	fmt.Println(fuscavo)

}
/*/

//Exercicio 4 - Struct
/*
package main

import (
	"fmt"
)

func main() {
	x := struct {
		nome string
		sabor string
		ondetem []string
		vaibemcom map[string]string
	}{
		nome: "Bolo",
		sabor: "doce",
		ondetem: []string{"no brasil", "nos EUA"},
		vaibemcom: map[string]string{
			"no café da manhã" : "café",
			"na café da tarde" : "chá",
		},
	}

	fmt.Println(x)
}
/*/

//Aula 87 - Funções
//Função: abstração e reutilização
/*
package main

import (
	"fmt"
)

func main() {
	s := "abc"

	fmt.Println(len(s))
}

func (receiver) nome (quantasletrastem) (int, string) {
}

func basica () {
	fmt.Println("Olá, bom dia!")
}

func soma (x, y int) int {
	return x + y
}

func soma (x int, y string)

func soma (x ,...int) int, int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
/*/

//Aula 88 - Enumerando uma slice
/*
package main

import (
	"fmt"
)

func main() {
	si := []int{10,10,10,1,2,3,4,5}

	total := soma(si...)

	fmt.Println(total)
}

func soma(x ...int) int{
	soma := 0

	for _, v := range x {
		soma += v
	}
	return soma
}
/*/

//Aula 89 - Defer
/*
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("com defer, veio primeiro")
	fmt.Println("sem defer, veio em segundo")
}
/*/

//Aula 90 - Metodos
/*
package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int
}

func (p pessoa)	oibomdia() {
	fmt.Println(p.nome, "diz bom dia")
}

func main() {
	mauricio := pessoa{"Mauricio", 30}
	mauricio.oibomdia()
}
/*/

//Aula 91 - Interfaces & Polimorfismos
/*
package main

import (
"fmt"
"math"
)

// Interface Forma
type Forma interface {
Area() float64
}

// Struct Retangulo
type Retangulo struct {
Largura, Altura float64
}

// Implementação do método Area para Retangulo
func (r Retangulo) Area() float64 {
return r.Largura * r.Altura
}

// Struct Círculo
type Circulo struct {
Raio float64
}

// Implementação do método Area para Circulo
func (c Circulo) Area() float64 {
return math.Pi * c.Raio * c.Raio
}

// Função que recebe qualquer Forma e imprime a área
func mostrarArea(f Forma) {
fmt.Printf("Área: %.2f\n", f.Area())
}

func main() {
r := Retangulo{Largura: 10, Altura: 5}
c := Circulo{Raio: 7}

mostrarArea(r) // Área do retângulo
mostrarArea(c) // Área do círculo
}
/*/

//Aula 92 - Funções Anonimas
/*
package main

import "fmt"

func main() {
// Função anônima atribuída a uma variável
saudacao := func(nome string) {
fmt.Printf("Olá, %s!\n", nome)
}

// Chamando a função
saudacao("Maria")

// Função anônima autoexecutável
func(a, b int) {
fmt.Printf("A soma de %d e %d é %d\n", a, b, a+b)
}(3, 4)
}
/*/

//Aula 96 - Closure
/*
package main

import "fmt"

// Função que retorna um closure
func contador() func() int {
contadorInterno := 0
return func() int {
contadorInterno++
return contadorInterno
}
}

func main() {
// Criando dois closures independentes
cont1 := contador()
cont2 := contador()

fmt.Println(cont1()) // 1
fmt.Println(cont1()) // 2
fmt.Println(cont2()) // 1
fmt.Println(cont1()) // 3
fmt.Println(cont2()) // 2
}
/*/

//Aula 97 - Recursividade
/*
package main

import "fmt"

// Função recursiva para calcular o fatorial
func fatorial(n int) int {
if n == 0 {
return 1
}
return n * fatorial(n-1)
}

func main() {
num := 5
fmt.Printf("Fatorial de %d é %d\n", num, fatorial(num))
}
/*/

//Aula 98 - Exercicios de Fixação
// Exercicio - 1
/*
package main

import (
"fmt"
)

func soma ()int {
		return 25
	}
func somador () (int, string){
		return 10, "dez"
	}

func main() {

	fmt.Println(soma())
	fmt.Println(somador())
}
/*/

//Exercicio 2
/*
package main

import (
	"fmt"
)

func soma(numeros ...int)int{
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	numero := []int{5, 10, 15}
	fmt.Println(soma(numero...))
}
/*/

// Exercicio 3
/*
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Tertuliano")
	fmt.Println("Thiago")
}
/*/

// Exercicio 4
/*
package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) nomecompleto(){
	fmt.Println(p.nome, p.sobrenome)
}

func main() {
	Thiago := pessoa{"Thiago","Tertuliano",20}
	Thiago.nomecompleto()
}
/*/

//Exercicio 5
/*
package main

import (
"fmt"
"math"
)

type Forma interface {
Area() float64
}

type Retangulo struct {
Largura, Altura float64
}

func (r Retangulo) Area() float64 {
return r.Largura * r.Altura
}

type Circulo struct {
Raio float64
}

func (c Circulo) Area() float64 {
return math.Pi * c.Raio * c.Raio
}

func mostrarArea(f Forma) {
fmt.Printf("Área: %.2f\n", f.Area())
}

func main() {
r := Retangulo{Largura: 10, Altura: 5}
c := Circulo{Raio: 7}

mostrarArea(r)
mostrarArea(c)
}
/*/

//Exercicio 6
/*
package main

import (
	"fmt"
)

func main() {
	gol := func(acao string){
		fmt.Printf("Olha o %s", acao)
	}

	gol("Golaço")
}
/*/

//Exercicio 7
/*
package main

import (
"fmt"
)

func main() {

numeros := []int{10, 20}

soma := func(nums []int) int {
total := 0
for _, n := range nums {
total += n
}
return total
}

resultado := soma(numeros)
fmt.Printf("A soma é: %d\n", resultado)
}
/*/

// Exercicio 8
/*
package main

import (
	"fmt"
)

func main() {
	numeros := []int{7, 7}

	multiplicador := func(numero []int) int {
		total := 1 //Se deixar zero em contas de multiplicação dará errado!
		for _, valor := range numero {
			total *= valor
		}
		return total
	}

	resultado := multiplicador(numeros)
	fmt.Println(resultado)
}
/*/

//Exercicio 9
/*
package main

import (
	"fmt"
)

func executarOperacao(a, b int, operacao func(int, int) int) {
	resultado := operacao(a, b)
	fmt.Println("Resultado:", resultado)
	}

func main() {

	soma := func(x, y int) int {
	return x + y
	}

	multiplicacao := func(x, y int) int {
	return x * y
	}

	executarOperacao(3, 4, soma)          // Resultado: 7
	executarOperacao(3, 4, multiplicacao) // Resultado: 12
}
/*/

//Exercicio 10
/*
package main

import (
	"fmt"
)

func contador() func() int {
	contadorExterno := 0
	return func() int {
		contadorExterno++
		return contadorExterno
	}
}

func main() {
	cont1 := contador()
	cont2 := contador()

	fmt.Println(cont1())
	fmt.Println(cont2())
}
/*/

// Aula 109 - O que são ponteiro?
//Um ponteiro é uma variável que armazena o endereço de memória de outra variável. Com ponteiros, você pode modificar o valor original de uma variável dentro de uma função, por exemplo.
/*
package main

import "fmt"

func dobrarNumero(n *int) {
	*n = *n * 2
}

func main() {
	num := 10
	fmt.Println("Antes:", num)

	dobrarNumero(&num)

	fmt.Println("Depois:", num)
}
// Oque esse código fez?
// n *int: a função recebe um ponteiro para int.
*n: acessa o valor armazenado no endereço de memória.
&num: passa o endereço da variável num para a função.
O valor de num é alterado diretamente na memória, sem precisar retornar nada.
//Aula 110 - Quando não usar ponteiros?
Ótima pergunta! Saber **quando usar ponteiros** em Go é essencial para escrever código mais eficiente e claro. Aqui vão os principais casos em que **vale a pena usar ponteiros**:

---

### ✅ **1. Quando você quer modificar o valor original de uma variável**
Se você passar uma variável para uma função **por valor**, ela será copiada. Com ponteiros, você pode alterar o valor original diretamente.

```go
func alterarNome(nome *string) {
	*nome = "Novo Nome"
}
```

---

### ✅ **2. Quando você quer evitar cópias desnecessárias**
Se você estiver lidando com **structs grandes**, passar por valor pode ser custoso. Usar ponteiros evita a cópia e melhora a performance.

```go
type Pessoa struct {
	Nome  string
	Idade int
	// imagine muitos outros campos...
}

func atualizarIdade(p *Pessoa) {
	p.Idade++
}
```

---

### ✅ **3. Quando você precisa representar "ausência de valor"**
Ponteiros podem ser `nil`, o que é útil para indicar que algo **não foi inicializado** ou está **opcional**.

```go
var p *Pessoa // nil
if p == nil {
	fmt.Println("Pessoa não definida")
}
```

---

### ✅ **4. Quando você quer compartilhar dados entre funções**
Se várias funções precisam acessar e modificar os mesmos dados, usar ponteiros garante que todas estejam trabalhando com a **mesma instância**.

---

### ❌ **Quando *não* usar ponteiros**
- Para **tipos primitivos pequenos** (como `int`, `bool`, `float64`) em funções simples, geralmente não vale a pena.
- Quando você **não precisa modificar** o valor original nem economizar memória.
/*/

//Aula 111 - Exercicios de Fixação
// Exercicio 1
/*
package main

import (
	"fmt"
)

func main() {

	valor := 42
	fmt.Println("Valor:", valor)
	fmt.Println("Endereço de memória:", &valor)
}
/*/
//Exercicio 2
/*
package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Caio"
	p.sobrenome = "Ribeiro"
}

func main() {
	thiago := pessoa{"Thiago", "Matos", 20}
	fmt.Println(thiago)
	mudeMe(&thiago)
	fmt.Println(thiago)
}
/*/

//Aula 113 - Documentação JSON
/*
package main

import (
"encoding/json"
"fmt"
)

// Pessoa representa uma pessoa com nome e idade.
// As tags `json:"..."` definem como os campos serão nomeados no JSON.
type Pessoa struct {
Nome  string `json:"nome"`  // Nome da pessoa
Idade int    `json:"idade"` // Idade da pessoa
}

func main() {
// Criando uma instância da struct Pessoa
p := Pessoa{
Nome:  "Thiago",
Idade: 30,
}

// Convertendo a struct para JSON
jsonData, err := json.Marshal(p)
if err != nil {
fmt.Println("Erro ao converter para JSON:", err)
return
}

// Exibindo o JSON gerado
fmt.Println(string(jsonData))

// Simulando recebimento de JSON e convertendo de volta para struct
jsonRecebido := `{"nome":"Renata","idade":25}`
var novaPessoa Pessoa

err = json.Unmarshal([]byte(jsonRecebido), &novaPessoa)
if err != nil {
fmt.Println("Erro ao converter JSON para struct:", err)
return
}

// Exibindo a struct preenchida a partir do JSON
fmt.Println(novaPessoa)
}
/*/

//Aula 114 - JSON Marchal (ordenação)
/*
package main

import (
"encoding/json"
"fmt"
"sort"
)

// Produto representa um item com nome, preço e quantidade.
// As tags JSON definem os nomes dos campos no JSON.
type Produto struct {
Nome      string  `json:"nome"`
Preco     float64 `json:"preco"`
Quantidade int    `json:"quantidade"`
}

func main() {
// Criando uma slice de produtos
produtos := []Produto{
{"Caderno", 12.90, 5},
{"Caneta", 2.50, 10},
{"Borracha", 1.20, 8},
}

// Ordenando os produtos por nome (ordem alfabética)
sort.Slice(produtos, func(i, j int) bool {
return produtos[i].Nome < produtos[j].Nome
})

// Convertendo a slice ordenada para JSON
jsonData, err := json.MarshalIndent(produtos, "", "  ")
if err != nil {
fmt.Println("Erro ao converter para JSON:", err)
return
}

// Exibindo o JSON formatado
fmt.Println(string(jsonData))
}

/*/

//Aula 115 - JSON Unmarshal (Desordenação)
/*
package main

import (
"encoding/json"
"fmt"
)

// Produto representa um item com nome, preço e quantidade.
// As tags JSON indicam como os campos devem ser mapeados.
type Produto struct {
Nome       string  `json:"nome"`
Preco      float64 `json:"preco"`
Quantidade int     `json:"quantidade"`
}

func main() {
// JSON com os campos em ordem diferente da struct
jsonRecebido := `{
"quantidade": 3,
"nome": "Lápis",
"preco": 1.75
}`

var p Produto

// Convertendo o JSON para a struct Produto
err := json.Unmarshal([]byte(jsonRecebido), &p)
if err != nil {
fmt.Println("Erro ao fazer unmarshal:", err)
return
}

// Exibindo a struct preenchida
fmt.Println("Struct preenchida a partir do JSON:")
fmt.Printf("Nome: %s\n", p.Nome)
fmt.Printf("Preço: %.2f\n", p.Preco)
fmt.Printf("Quantidade: %d\n", p.Quantidade)
}
/*/

//Aula 116 - Interface Writer
/*
package main

import (
"fmt"
"io"
"os"
"strings"
)

func main() {
// Criando uma string que será usada como fonte de dados
texto := "Olá, mundo! Este é um exemplo da interface Writer."

// Criando um Writer que escreve no terminal (stdout)
var escritor io.Writer = os.Stdout

// Escrevendo diretamente no terminal
// A função Write espera um slice de bytes ([]byte)
_, err := escritor.Write([]byte(texto))
if err != nil {
fmt.Println("Erro ao escrever:", err)
}

fmt.Println("\n---")

// Criando outro Writer, agora usando um buffer de string
var buffer strings.Builder

// strings.Builder também implementa a interface io.Writer
_, err = buffer.Write([]byte("Texto armazenado no buffer."))
if err != nil {
fmt.Println("Erro ao escrever no buffer:", err)
}

// Exibindo o conteúdo do buffer
fmt.Println(buffer.String())
}
/*/

//Aula 117 - Pacote Sort
/*
package main

import (
"fmt"
"sort"
)

func main() {
// Slice de inteiros desordenado
numeros := []int{42, 7, 19, 3, 99}

// Ordenando o slice de inteiros em ordem crescente
sort.Ints(numeros)

// Exibindo o resultado
fmt.Println("Slice ordenado:", numeros)

// Slice de strings
nomes := []string{"Carlos", "Ana", "Bruno", "Daniela"}

// Ordenando o slice de strings em ordem alfabética
sort.Strings(nomes)

fmt.Println("Nomes ordenados:", nomes)

// Slice de structs com ordenação personalizada
type Pessoa struct {
Nome string
Idade int
}

pessoas := []Pessoa{
{"João", 30},
{"Maria", 25},
{"Pedro", 40},
}

// Ordenando por idade (do mais novo para o mais velho)
sort.Slice(pessoas, func(i, j int) bool {
return pessoas[i].Idade < pessoas[j].Idade
})

fmt.Println("Pessoas ordenadas por idade:")
for _, p := range pessoas {
fmt.Printf("%s - %d anos\n", p.Nome, p.Idade)
}
}
/*/

//Aula 118 - Customizando Sort
/*
package main

import (
"fmt"
"sort"
)

// Definindo uma struct Pessoa
type Pessoa struct {
Nome  string
Idade int
Cidade string
}

func main() {
// Slice de pessoas
pessoas := []Pessoa{
{"João", 30, "São Paulo"},
{"Maria", 25, "Rio de Janeiro"},
{"Pedro", 25, "Belo Horizonte"},
{"Ana", 40, "Curitiba"},
}

// Ordenando por idade (crescente) e, em caso de empate, por nome (alfabético)
sort.Slice(pessoas, func(i, j int) bool {
if pessoas[i].Idade == pessoas[j].Idade {
return pessoas[i].Nome < pessoas[j].Nome
}
return pessoas[i].Idade < pessoas[j].Idade
})

// Exibindo o resultado
fmt.Println("Pessoas ordenadas por idade e nome:")
for _, p := range pessoas {
fmt.Printf("%s - %d anos - %s\n", p.Nome, p.Idade, p.Cidade)
}
}
/*/

//Aula 119 - Bcrypt
/*
package main

import (
"fmt"
"golang.org/x/crypto/bcrypt" //Instalar bycript: go get golang.org/x/crypto/bcrypt
)

func main() {
// Senha original (texto plano)
senha := "minhaSenha123"

// Gerando o hash da senha com custo padrão (recomendado: 10 ou mais)
hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
if err != nil {
fmt.Println("Erro ao gerar hash:", err)
return
}

// Exibindo o hash gerado (é diferente a cada execução)
fmt.Println("Hash gerado:", string(hash))

// Verificando se a senha fornecida bate com o hash
err = bcrypt.CompareHashAndPassword(hash, []byte(senha))
if err != nil {
fmt.Println("Senha incorreta!")
} else {
fmt.Println("Senha correta!")
}
}
/*/

// Aula 119 - Exercicios de Fixação
/*
package main

import (
"encoding/json"
"fmt"
"sort"
)

type Clube struct {
Nome       string `json:"nome"`
Modalidade string `json:"modalidade"`
Fundacao   string `json:"fundacao"`
Idade      int    `json:"idade"`
}

func main() {

clubes := []Clube{
{"São Paulo", "Futebol", "10/10/1930", 95},
{"Corinthians", "Futebol", "10/10/1910", 115},
{"Palmeiras", "Futebol", "10/10/1915", 110},
}

sort.Slice(clubes, func(i, j int) bool {
return clubes[i].Nome < clubes[j].Nome
})

jsonData, err := json.MarshalIndent(clubes, "", "  ")
if err != nil {
fmt.Println("Erro ao converter para JSON:", err)
return
}

fmt.Println(string(jsonData))
}
/*/

//Exercicio 2
/*
package main

import (
"encoding/json"
"fmt"
)

type Jogador struct {
Nome          string `json:"nome"`
NumeroCamisa  int    `json:"numeroCamisa"`
Posicao       string `json:"posicao"`
}

func main() {
jsonRecebido := `{
"nome": "Thiago",
"numeroCamisa": 10,
"posicao": "Atacante"
}`

var j Jogador

err := json.Unmarshal([]byte(jsonRecebido), &j)
if err != nil {
fmt.Println("Erro ao fazer o unmarshal:", err)
return
}

fmt.Println("Struct preenchida a partir do JSON:")
fmt.Printf("Nome do jogador: %s\n", j.Nome)
fmt.Printf("Número da camisa: %d\n", j.NumeroCamisa)
fmt.Printf("Posição: %s\n", j.Posicao)
}
/*/

//Exercicio 3
/*
package main

import (
"encoding/json"
"fmt"
"os"
)

type Jogador struct {
Nome         string `json:"nome"`
NumeroCamisa int    `json:"numeroCamisa"`
Posicao      string `json:"posicao"`
}

func main() {
j := Jogador{
Nome:         "Thiago",
NumeroCamisa: 10,
Posicao:      "Atacante",
}

fmt.Println("JSON gerado com json.NewEncoder(os.Stdout):")

encoder := json.NewEncoder(os.Stdout)

err := encoder.Encode(j)
if err != nil {
fmt.Println("Erro ao codificar JSON:", err)
}
}
/*/

//Exercicio 4
/*
package main

import (
"encoding/json"
"fmt"
"os"
"sort"
)

type DadosOrdenados struct {
Numeros []int    `json:"numeros"`
Nomes   []string `json:"nomes"`
}

func main() {
numeros := []int{42, 7, 19, 3, 99}
nomes := []string{"Carlos", "Ana", "Bruno", "Daniela"}

sort.Ints(numeros)
sort.Strings(nomes)

dados := DadosOrdenados{
Numeros: numeros,
Nomes:   nomes,
}

fmt.Println("JSON com slices ordenados:")

encoder := json.NewEncoder(os.Stdout)
encoder.SetIndent("", "  ") // Formatação bonita

err := encoder.Encode(dados)
if err != nil {
fmt.Println("Erro ao codificar JSON:", err)
}
}
/*/

//Exercicio 5
/*
package main

import (
"encoding/json"
"fmt"
"os"
"sort"
)

type User struct {
Nome      string `json:"nome"`
Sobrenome string `json:"sobrenome"`
Idade     int    `json:"idade"`
}

func main() {
usuarios := []User{
{"Ana", "Silva", 30},
{"Carlos", "Oliveira", 25},
{"Bruno", "Almeida", 25},
{"Daniela", "Silva", 40},
{"Ana", "Costa", 30},
}

sort.Slice(usuarios, func(i, j int) bool {
if usuarios[i].Idade == usuarios[j].Idade {
return usuarios[i].Sobrenome < usuarios[j].Sobrenome
}
return usuarios[i].Idade < usuarios[j].Idade
})

fmt.Println("Usuários ordenados por idade e sobrenome:")

encoder := json.NewEncoder(os.Stdout)
encoder.SetIndent("", "  ")

err := encoder.Encode(usuarios)
if err != nil {
fmt.Println("Erro ao codificar JSON:", err)
}
}
/*/

//Aula 126 - Goroutines e WaitGroups
/*
package main

import (
"fmt"
"sync"
"time"
)

// Função que simula uma tarefa demorada
func tarefa(id int, wg *sync.WaitGroup) {
defer wg.Done() // Marca que a goroutine terminou

fmt.Printf("Tarefa %d iniciada\n", id)
time.Sleep(2 * time.Second) // Simula trabalho
fmt.Printf("Tarefa %d finalizada\n", id)
}

func main() {
var wg sync.WaitGroup // Criando um WaitGroup

// Iniciando 3 goroutines
for i := 1; i <= 3; i++ {
wg.Add(1) // Informando que uma nova goroutine será iniciada
go tarefa(i, &wg)
}

// Espera todas as goroutines terminarem
wg.Wait()

fmt.Println("Todas as tarefas foram concluídas.")
}
/*/
//Aula 128 - Condição de Corrida
/*
package main

import (
"fmt"
"time"
)

var contador int

func incrementar() {
for i := 0; i < 1000; i++ {
contador++
}
}

func main() {
for i := 0; i < 10; i++ {
go incrementar()
}

time.Sleep(2 * time.Second)
fmt.Println("Valor final do contador:", contador)
}
/*/

//Aula 129 - Mutex
/*
package main

import (
	"fmt"
	"time"
)

var contador int

func incrementar() {
	for i := 0; i < 1000; i++ {
		contador++
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go incrementar()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Valor final do contador:", contador)
}
/*/

//Aula 130 - Atomic

/*
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var counter int64 = 0
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            atomic.AddInt64(&counter, 1)
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Contador final:", counter)
}
/*/

//Aula 139 - Exercicios de Fixação
//Exercicio 1
/*
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novasgoroutines(100)
	wg.Wait()
}

func novasgoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número:", i+1)
			wg.Done()
		}(x)
	}
}

// - duas goroutines adicionais
// - cada uma faz um print
// - waitgroups
/*/

//Exercicio 2
/*
package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz oi!")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	p1 := pessoa{"Genghis Khan", 1000}

	p1.falar() // ← É um shortcut pra (&p1).falar()

	(&p1).falar() // ← É a maneira "mais correta."

	dizerAlgumaCoisa(&p1) // ← Funciona!

	// dizerAlgumaCoisa(p1) // ← Não funciona!
}
/*/

//Exercicio 3
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const quantidadeDeGoroutines = 50000

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}

// - Utilizando goroutines, crie um programa incrementador:
// - Tenha uma variável com o valor da contagem
// - Crie um monte de goroutines, onde cada uma deve:
// 	- Ler o valor do contador
// 	- Salvar este valor
// 	- Fazer yield da thread com runtime.Gosched()
// 	- Incrementar o valor salvo
// 	- Copiar o novo valor para a variável inicial
// - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
// - Demonstre que há uma condição de corrida utilizando a flag -race
/*/

//Exercicio 4
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

const quantidadeDeGoroutines = 50000

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}

// - Utilizando goroutines, crie um programa incrementador:
// - Tenha uma variável com o valor da contagem
// - Crie um monte de goroutines, onde cada uma deve:
// 	- Ler o valor do contador
// 	- Salvar este valor
// 	- Fazer yield da thread com runtime.Gosched()
// 	- Incrementar o valor salvo
// 	- Copiar o novo valor para a variável inicial
// - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
// - Demonstre que há uma condição de corrida utilizando a flag -race
/*/

//Exercicio 5
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

const quantidadeDeGoroutines = 50000

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			// func AddInt32(addr *int32, delta int32) (new int32)
			// func LoadInt32(addr *int32) (val int32)
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}
/*/

//Exercicio 6

/*
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
}

// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//     - go run
//     - go build
//     - go install
/*/

//Aula 146 - Canais
/*
package main

import (
    "fmt"
)

func sayHello(ch chan string) {
    ch <- "Olá de uma goroutine!"
}

func main() {
    ch := make(chan string)

    go sayHello(ch)

    msg := <-ch
    fmt.Println(msg)
}
/*/

//Aula 147 - Canais direcionais
/*
package main

import "fmt"

func enviar(ch chan<- string) {
    ch <- "Olá de enviar()"
}

func receber(ch <-chan string) {
    msg := <-ch
    fmt.Println("Recebido:", msg)
}

func main() {
    ch := make(chan string)

    go enviar(ch)
    receber(ch)
}
/*/

//Aula 148 - Range e close
/*
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // importante!
	}()

	for val := range ch {
		fmt.Println("Recebido:", val)
	}
}
/*/

//Aula 149 - Select
/*
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Mensagem de ch1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Mensagem de ch2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Recebido:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Recebido:", msg2)
        }
    }
}
/*/

//Aula 150 - Expressão comma-ok
/*
package main

import "fmt"

func main() {
    ch := make(chan int)
    go func() {
        ch <- 10
        close(ch)
    }()

    val, ok := <-ch
    fmt.Println("Valor:", val, "OK:", ok)

    val, ok = <-ch
    fmt.Println("Valor:", val, "OK:", ok)
}
/*/

//Aula 151 - Convergência
/*
package main

import (
    "fmt"
    "time"
)

func produtor(id int, out chan<- string) {
    for i := 0; i < 3; i++ {
        msg := fmt.Sprintf("Produtor %d: %d", id, i)
        out <- msg
        time.Sleep(time.Millisecond * 100)
    }
}

func main() {
    saida := make(chan string)

    // Iniciando múltiplos produtores
    for i := 1; i <= 3; i++ {
        go produtor(i, saida)
    }

    // Recebendo 9 mensagens (3 de cada produtor)
    for i := 0; i < 9; i++ {
        fmt.Println(<-saida)
    }
}
/*/

//Aula 152 - Divergência
/*
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processando job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	// Criando 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Enviando 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // importante para encerrar os workers

	wg.Wait()
}
/*/

//Aula 153 - Context
/*
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelado\n", id)
            return
        default:
            fmt.Printf("Worker %d trabalhando...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Cancelando todos os workers...")
    cancel()

    time.Sleep(1 * time.Second)
}
/*/

//Aula 154 - Exercicios de Fixação
//Exercicio 1
/*
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
/*/

//Exercicio 2
/*
package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
/*/

//Exercicio 3
/*
package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(canal <-chan int) {
	for v := range canal {
		fmt.Println(v)
	}
}
/*/

//Exercicio 4
/*
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

func receive (c <-chan int, q chan int) {
	for {
		select {
			case v := <-c:
				fmt.Println(v)
			case <-q:
				return
		}
	}
}
/*/

//Exercicio 5
/*
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
/*/

//Exercicio 6

// - Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
/*
package main

import "fmt"

func main() {
	c := make(chan int)
	go botalá(c)
	for v := range c {
		fmt.Println(v)
	}
}

func botalá(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
/*/

//Exercicio 7

// - Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.
/*
package main

import "fmt"

func main() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, "\t", <-canal)
	}
}
/*/

//Aula 161 - Tratamento de Erros

/*🧠 1. Entendendo Erros em Go

Go trata erros de forma diferente de linguagens como Python ou Java. Em vez de usar exceções, Go utiliza valores de erro explícitos que são retornados pelas funções.

📌 Como os erros funcionam em Go?

Em Go, uma função que pode falhar geralmente retorna dois valores:

resultado, err := minhaFuncao()
if err != nil {
    // tratar o erro
}


🔍 Tipos Comuns de Erros em Go

1. Erros de Compilação
   - Ocorrem quando há problemas de sintaxe ou tipos incompatíveis.
   - Exemplo:
     ```go
     var x int = "texto" // erro: tipos incompatíveis
     ```

2. Erros de Execução (Runtime)
   - Como divisão por zero ou acesso a ponteiros nulos.
   - Exemplo:
     var p *int
     fmt.Println(*p) // panic: invalid memory address

3. Erros Retornados por Funções
   - A forma mais comum de lidar com erros em Go.
   - Exemplo:
     package main

     import (
         "fmt"
         "os"
     )

     func main() {
         file, err := os.Open("arquivo.txt")
         if err != nil {
             fmt.Println("Erro ao abrir o arquivo:", err)
             return
         }
         defer file.Close()
     }

🧰 O tipo `error`

O tipo `error` é uma interface embutida em Go:


type error interface {
    Error() string
}

Você pode criar seus próprios erros com `errors.New` ou `fmt.Errorf`:

import "errors"

func minhaFuncao() error {
    return errors.New("algo deu errado")
}

// Aula 162 - Verificando Erros

Em Go, **verificar erros** é uma prática essencial. Como vimos, funções que podem falhar geralmente retornam um valor do tipo `error`. A verificação é feita com uma simples comparação:

if err != nil {
    // tratar o erro
}

📌 Exemplo Básico

package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("arquivo_inexistente.txt")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close()
}

Neste exemplo:
- A função `os.Open` retorna um `*File` e um `error`.
- Se o arquivo não existir, `err` será diferente de `nil`.
- O erro é verificado com `if err != nil`.


🧠 Boas Práticas ao Verificar Erros

1. Sempre verifique o erro imediatamente após a chamada da função.
   resultado, err := fazerAlgo()
   if err != nil {
       // trate o erro aqui
   }

2. Forneça mensagens úteis.
   if err != nil {
       log.Fatalf("Erro ao processar dados: %v", err)
   }

3. Use `errors.Is` ou `errors.As` para verificar tipos específicos de erro (Go 1.13+).

   if errors.Is(err, os.ErrNotExist) {
       fmt.Println("Arquivo não existe.")
   }

🧪 Exemplo com `errors.Is`

package main

import (
    "errors"
    "fmt"
    "os"
)

func main() {
    _, err := os.Open("arquivo.txt")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("Arquivo não encontrado.")
    } else if err != nil {
        fmt.Println("Outro erro:", err)
    }
}

🧰 Dica Extra: Ignorando erros (com cuidado!)

Você pode ignorar um erro usando `_`, mas isso **não é recomendado** a menos que tenha certeza de que o erro pode ser ignorado:

_, _ = fmt.Println("Olá") // ignora o erro de retorno

// Aula 163 - Print & Log

🧾 1. Usando `fmt.Print` para erros

O pacote `fmt` é frequentemente usado para imprimir mensagens no terminal, incluindo erros:

import "fmt"

err := algumaFuncao()
if err != nil {
    fmt.Println("Erro:", err)
}

✅ Vantagens:
- Simples e direto.
- Útil para testes rápidos ou scripts pequenos.

⚠️ Limitações:
- Não inclui informações como **timestamp**.
- Não é ideal para **aplicações maiores ou produção**.


🧰 2. Usando o pacote `log`

O pacote `log` oferece uma forma mais robusta de registrar erros e mensagens:

import "log"

err := algumaFuncao()
if err != nil {
    log.Println("Erro:", err)
}

🧠 Recursos do `log`:
- Adiciona **data e hora** automaticamente.
- Pode ser redirecionado para arquivos ou sistemas de log.
- Tem funções como:
  - `log.Print()`
  - `log.Println()`
  - `log.Fatalf()` → imprime e encerra o programa
  - `log.Panicf()` → imprime e gera um panic

🧪 Exemplo comparativo
package main

import (
    "errors"
    "fmt"
    "log"
)

func retornaErro() error {
    return errors.New("algo deu errado")
}

func main() {
    err := retornaErro()

    if err != nil {
        fmt.Println("Usando fmt:", err)
        log.Println("Usando log:", err)
    }
}

//Aula 164 - Recover

package main

import "fmt"

func proteger() {
    if r := recover(); r != nil {
        fmt.Println("Recuperado de um panic:", r)
    }
}

func podeFalhar() {
    defer proteger()
    panic("algo deu muito errado!")
}

func main() {
    podeFalhar()
    fmt.Println("Programa continua executando normalmente.")
}

// Aula 165 - Erros com informações adicionais

package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("registro não encontrado")

func buscarRegistro(id int) error {
    if id == 0 {
        return fmt.Errorf("ID inválido (%d): %w", id, ErrNotFound)
    }
    return nil
}

func main() {
    err := buscarRegistro(0)
    if err != nil {
        fmt.Println("Erro:", err)

        if errors.Is(err, ErrNotFound) {
            fmt.Println("Tratamento específico para registro não encontrado")
        }
    }
}

//Aula 166 - Exercícios de Fixação
// Exercício 1
/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))

}
/*/

// Exercício 2
/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("bino, é uma cilada!")
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("é uma cilada bino!")
	}
	return bs, nil
}
/*/

// Exercício 3
/*
package main

import "fmt"

// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
// - Crie uma função que tenha um valor do tipo error como parâmetro.
// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.

type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("deu zica! e olha isso: %v", e.qualquercoisa)
}

func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(erroEspecial).qualquercoisa, "\nno método Error, eu tenho:", e)
}

func main() {
	arg := erroEspecial{"EMOJIS!!!!!!!!"}
	erroComoParametro(arg)
}
/*/

// Exercício 4
/*
package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		erroNovo := fmt.Errorf("Deu erro no valor: %v", f)
		return 0, sqrtError{"123", "321", erroNovo}
	}
	return 42, nil
}
/*/

// Aula 172 - Documentação  Go Doc
/*
📘 1. O que é `go doc`?

`go doc` é um comando da ferramenta `go` que permite acessar a documentação de qualquer item do seu código ou da biblioteca padrão, sem sair do terminal.

2. Como usar `go doc`**

Documentação de um pacote

go doc fmt

Mostra a documentação do pacote `fmt`.

Função ou tipo específico

go doc fmt.Println

Mostra a documentação da função `Println` do pacote `fmt`.

Tipo e seus métodos

go doc net/http Client

Mostra a documentação do tipo `Client` no pacote `net/http`.

Método de um tipo

go doc net/http Client.Do

Mostra a documentação do método `Do` do tipo `Client`.

🧪 3. Exemplo prático

go doc encoding/json Marshal


Saída esperada:

func Marshal(v interface{}) ([]byte, error)
    Marshal retorna a codificação JSON de v.

🧠 4. Dicas úteis

- Use `go doc` dentro de um projeto Go para consultar pacotes locais.
- Combine com `go list` para descobrir nomes de pacotes:
  
  go list ./...

🧩 Resumo

|         Comando             |         O que faz          |
|-----------------------------|----------------------------|
|     `go doc pacote`         |     Mostra a doc do pacote |
|     `go doc pacote.Função`  | Mostra a doc de uma função |
|     `go doc pacote.Tipo`    |    Mostra a doc de um tipo |
| `go doc pacote.Tipo.Método` |  Mostra a doc de um método |

/*

// Aula 173 - Go doc
/*
📚 1. O que é `godoc`?

`godoc` é uma ferramenta que:

- Gera e exibe a **documentação de pacotes Go**.
- Pode ser usada tanto no terminal quanto como um **servidor web local**.
- Lê os comentários do código-fonte e os transforma em documentação navegável.

🧰 2. Como usar `godoc`

Instalar (se necessário)

Se não estiver disponível, instale com:


go install golang.org/x/tools/cmd/godoc@latest

Rodar como servidor local


godoc -http=:6060

- Isso inicia um servidor web local em `http://localhost:6060`.
- Você pode navegar pela documentação da biblioteca padrão e dos seus próprios pacotes.

Usar no terminal

godoc fmt Println

- Mostra a documentação da função `Println` do pacote `fmt`.

🌐 3. Navegando no navegador

Acesse:

http://localhost:6060/pkg/

Você verá uma interface parecida com a documentação oficial do Go, mas gerada localmente com base no seu ambiente.

✍️ 4. Escrevendo boa documentação para `godoc`

- Comentários acima de funções, tipos e variáveis **devem começar com o nome do item**:

// Soma retorna a soma de dois inteiros.
func Soma(a, b int) int {
    return a + b
}

- Isso garante que `godoc` reconheça e exiba corretamente a descrição.

🧩 Resumo

|           Comando            |           O que faz            |
|------------------------------|--------------------------------|
|       `godoc fmt Println`    | Mostra doc da função `Println` |
|       `godoc -http=:6060`    |   Inicia servidor web local    |
| `http://localhost:6060/pkg/` |  Interface web da documentação |


// Aula 174 - Godoc.org
/*
🌐 1. O que era o godoc.org?
	Era um site que exibia a documentação de pacotes Go diretamente do código-fonte hospedado em repositórios públicos (como GitHub).
	Permitía buscar por pacotes, tipos, funções e exemplos.
	Era alimentado automaticamente por repositórios públicos com código Go.
🔄 2. Substituição por pkg.go.dev
	Desde 2020, o site pkg.go.dev substituiu oficialmente o godoc.org.
📦 pkg.go.dev oferece:
	Interface moderna e responsiva.
	Suporte a múltiplas versões de pacotes.
	Integração com módulos Go (go mod).
	Exibição de exemplos, dependências e segurança.
🧭 3. Como usar pkg.go.dev (antigo godoc.org)
	Para ver a documentação de um pacote:

https://pkg.go.dev/github.com/usuário/repositorio
Exemplo:

https://pkg.go.dev/github.com/gorilla/mux
Você pode buscar diretamente na barra de pesquisa do site.

🧩 Resumo
|    Site    |      Status      |           Função                      |
| godoc.org  |	Descontinuado	|  Exibição de documentação pública     |
| pkg.go.dev |	   Ativo        |	Substituto oficial com mais recursos|

/*/

// Aula 175 - Escrevendo documentação
/*
✍️ 1. Comentários de documentação

Em Go, a documentação é escrita como comentários acima de funções, tipos, variáveis e pacotes. Para que ela seja reconhecida por ferramentas como `go doc`, o comentário deve:

- Estar imediatamente acima do item documentado.
- Começar com o nome do item.

🧪 Exemplo básico

// Soma retorna a soma de dois inteiros.
func Soma(a, b int) int {
    return a + b
}

📦 2. Documentando pacotes

No início do arquivo principal do pacote (`doc.go` ou `main.go`), você pode documentar o pacote inteiro:

// Pacote calculadora fornece funções matemáticas básicas.
package calculadora


🧰 3. Documentando tipos e métodos

// Pessoa representa um indivíduo com nome e idade.
type Pessoa struct {
    Nome string
    Idade int
}

// Saudacao retorna uma saudação personalizada para a pessoa.
func (p Pessoa) Saudacao() string {
    return "Olá, " + p.Nome
}

💡 4. Boas práticas

- Seja claro e direto.
- Use exemplos simples quando possível.
- Evite repetir o nome da função nos parâmetros.
- Use comentários consistentes em todo o projeto.

📄 5. Exemplo completo de documentação

// Pacote saudacao fornece funções para gerar mensagens de boas-vindas.
package saudacao

// BemVindo retorna uma mensagem de boas-vindas para o nome fornecido.
func BemVindo(nome string) string {
    return "Bem-vindo, " + nome + "!"
}

🧩 Resumo

|      Elemento    |                  Como documentar                        |
|------------------|---------------------------------------------------------|
|      Função      |    Comentário acima, começando com o nome da função     |
|      Tipo        | Comentário acima do `type`, explicando o que representa |
|      Método      |    Comentário acima, explicando o comportamento         |
|      Pacote      |    Comentário no topo do arquivo principal              |
/*/

// Aula 176 - Exercícios de Fixação
// Exercício 1
/*
// Arquivo principal: main.go
package main

import (
	"fmt"

	//"go/cachorro.go"
)

func main() {
	fmt.Println(cachorro.Idade(10))
}

// Arquivo cachorro.go
package cachorro 

import "fmt"

func Idade(anos int) string { 
	return fmt.Sprintf("A idade do cachorro é %d anos", anos)
}
/*/

// Aula 179 - Testes & Benchmarks – Introdução
/*
1. O que são testes em Go?

Go possui suporte nativo a testes através do pacote **`testing`**. Com ele, você pode:

- Escrever **testes unitários** (`TestXxx`)
- Criar **testes de benchmark** (`BenchmarkXxx`)
- Escrever **exemplos executáveis** (`ExampleXxx`)

🧪 2. Estrutura básica de um teste

Um arquivo de teste em Go:

- Deve ter o sufixo `_test.go`
- Deve importar o pacote `testing`
- Cada função de teste começa com `Test` e recebe um `*testing.T`

Exemplo:

// arquivo: calculo_test.go
package calculo

import "testing"

func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    esperado := 5

    if resultado != esperado {
        t.Errorf("esperado %d, mas obteve %d", esperado, resultado)
    }
}

🚀 3. Rodando testes

No terminal, dentro do diretório do projeto:

go test

Para ver mais detalhes:

go test -v

🏁 4. O que são benchmarks?

Benchmarks medem o **desempenho** de funções. São úteis para otimizar código.

Exemplo:

func BenchmarkSoma(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Soma(2, 3)
    }
}


E para rodar benchmarks:

go test -bench=.


🧩 Resumo

|     Tipo       |      Nome da função          |               Objetivo               |
|----------------|------------------------------|--------------------------------------|
| Teste unitário |   `TestXxx(t *testing.T)`    |     Verifica se o código funciona    |
|   Benchmark    | `BenchmarkXxx(b *testing.B)` |          Mede desempenho             |
|    Exemplo     |      `ExampleXxx()`          | Gera documentação e pode ser testado |

//Aula 180 - Testes em tabelas
/*

📋 1. O que são testes em tabela?

São testes que usam uma **estrutura de dados (geralmente um slice de structs)** para definir múltiplos cenários de teste. Cada entrada da tabela representa um **caso de teste** com entradas e saídas esperadas.


🧪 2. Exemplo prático

Vamos testar uma função `Soma(a, b int) int`:

func Soma(a, b int) int {
    return a + b
}

🧾 Teste em tabela:

package calculo

import "testing"

func TestSoma(t *testing.T) {
    testes := []struct {
        nome     string
        a, b     int
        esperado int
    }{
        {"soma positiva", 2, 3, 5},
        {"soma com zero", 0, 5, 5},
        {"soma negativa", -2, -3, -5},
        {"mistura", -1, 4, 3},
    }

    for _, tt := range testes {
        t.Run(tt.nome, func(t *testing.T) {
            resultado := Soma(tt.a, tt.b)
            if resultado != tt.esperado {
                t.Errorf("esperado %d, obteve %d", tt.esperado, resultado)
            }
        })
    }
}

3. Vantagens dos testes em tabela

- **Organização**: todos os casos ficam agrupados.
- **Facilidade de manutenção**: adicionar novos testes é simples.
- **Clareza**: cada teste tem um nome descritivo com `t.Run`.

🧩 Resumo

|     Elemento        |             Função            |
|---------------------|-------------------------------|
|   Slice de structs  |    Define os casos de teste   |
| `t.Run(nome, func)` |   Executa cada caso com nome  |
|     `t.Errorf`      |         Reporta falhas        |

// Aula 182 - go fmt, govet e golint 
/*

🧹 1. `go fmt` – Formatação automática

✅ O que faz:
- Formata seu código Go de acordo com o **padrão oficial da linguagem**.
- Corrige automaticamente espaçamentos, indentação, quebras de linha, etc.

📦 Como usar:

go fmt ./...


🧠 Por que usar:
- Garante **consistência visual** no código.
- Evita discussões sobre estilo em revisões de código.

🕵️ 2. `go vet` – Verificador estático

✅ O que faz:
- Analisa o código em busca de **erros sutis** que o compilador não detecta.
- Exemplo: chamadas incorretas de `fmt.Printf`, uso de variáveis não inicializadas, etc.

📦 Como usar:

go vet ./...


🧠 Por que usar:
- Ajuda a encontrar **bugs potenciais** antes da execução.
- É uma ferramenta de **análise estática** complementar ao compilador.


🧽 3. `golint` – Sugestões de estilo (opcional)

> ⚠️ `golint` está **obsoleto** e não é mais mantido oficialmente, mas ainda é usado em alguns projetos.

O que faz:
- Fornece **sugestões de estilo e boas práticas**.
- Exemplo: nomes de funções exportadas devem ter comentários iniciando com o nome da função.

📦 Como instalar:

go install golang.org/x/lint/golint@latest


📦 Como usar:

golint ./...


🧠 Por que usar:
- Ajuda a manter o código **idiomático** e alinhado com a comunidade Go.
- Melhora a **legibilidade** e a **manutenção** do código.

🧩 Resumo

| Ferramenta |       Função         |     Uso principal     |
|------------|----------------------|-----------------------|
| `go fmt`   |   Formata o código   | Estilo e legibilidade |
| `go vet`   | Verifica erros sutis | Qualidade e segurança |
| `golint`   | Sugere boas práticas | Estilo e documentação |

//