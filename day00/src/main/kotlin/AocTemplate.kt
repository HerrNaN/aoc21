import java.io.File

import cc.ekblad.konbini.*
import cc.ekblad.konbini.ParserResult.Error
import cc.ekblad.konbini.ParserResult.Ok
import cc.ekblad.konbini.parseToEnd

fun main() {
    val input = getInput()
    val answer = when (System.getenv("part")) {
        "part2" -> solutionPart2(input)
        else -> solutionPart1(input)
    }
    println(answer)
}

fun getInput(): List<Long> {
    val inputFile = File("input.txt").readText()
    val inputParser = parser{
        // Add day specific parser here!
        chain(integer, whitespace).terms
    }

    return when (val pRes = inputParser.parseToEnd(inputFile)) {
        is Ok -> pRes.result
        is Error -> throw Error("couldn't parse input: " + pRes.reason)
    }
}

fun solutionPart1(input: List<Long>) = input.sum()

fun solutionPart2(input: List<Long>) = input.reduce(Long::times)
