package main

import (
	"bytes"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Machine", func() {
	Describe("NewMachine", func() {
		It("returns a pointer to a new Machine instance", func() {
			m := NewMachine("code", os.Stdin, os.Stdout)
			Ω(*m).Should(BeAssignableToTypeOf(Machine{}))
		})
	})

	Describe("readChar", func() {
		It("reads into buf and stores the value in memory at current dp", func() {
			reader := strings.NewReader(">")
			m := &Machine{
				input: reader,
				buf:   make([]byte, 1),
				dp:    4,
			}
			m.readChar()
			Ω(m.buf[0]).Should(Equal(byte(62)))
			Ω(m.memory[4]).Should(Equal(62))
		})
	})

	Describe("putChar", func() {
		It("puts current value in memory at dp into buf and writes to output", func() {
			var b bytes.Buffer
			m := Machine{
				output: &b,
				buf:    make([]byte, 1),
				dp:     4,
			}
			m.memory[4] = 62
			m.putChar()
			Ω(b.Len()).Should(Equal(1))
			Ω(b.String()).Should(Equal(">"))
			Ω(b.Next(1)).Should(Equal(m.buf))
		})
	})

	Describe("Execute", func() {
		It("handles the '+' symbol by incrementing the value in memory at dp by 1", func() {
			m := Machine{}
			m.dp = 4
			m.memory[4] = 13
			m.code = "+"
			m.Execute()
			Ω(m.memory[4]).Should(Equal(14))
		})

		It("handles the '-' symbol by decrementing the value in memory at dp by 1", func() {
			m := Machine{}
			m.dp = 6
			m.memory[6] = 12
			m.code = "-"
			m.Execute()
			Ω(m.memory[6]).Should(Equal(11))
		})

		It("handles the > symbol by increasing dp by 1", func() {
			m := Machine{}
			m.dp = 19
			m.code = ">"
			m.Execute()
			Ω(m.dp).Should(Equal(20))
		})

		It("handles the < symbol by decreasing dp by 1", func() {
			m := Machine{}
			m.dp = 11
			m.code = "<"
			m.Execute()
			Ω(m.dp).Should(Equal(10))
		})

		It("handles the , symbol by calling readChar", func() {
			reader := strings.NewReader(".")
			m := Machine{}
			m.input = reader
			m.buf = make([]byte, 1)
			m.dp = 4
			m.code = ","
			m.Execute()
			Ω(m.memory[4]).Should(Equal(46))
		})

		It("handles the . symbol by calling putChar", func() {
			var b bytes.Buffer
			m := Machine{}
			m.output = &b
			m.buf = make([]byte, 1)
			m.dp = 4
			m.memory[4] = 60
			m.code = "."
			m.Execute()
			Ω(b.String()).Should(Equal("<"))
		})
	})
})
