import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

internal class AocTemplateTest {

    @Test
    internal fun `solutionPart1 should return sum of input`() {
        assertEquals(solutionPart1(listOf(listOf(1000,2000,3000), listOf(4000), listOf(5000,6000), listOf(7000,8000,9000), listOf(10000))), 24000)
    }

    @Test
    internal fun `solutionPart2 should return product of input`() {
        assertEquals(solutionPart2(listOf(listOf(1000,2000,3000), listOf(4000), listOf(5000,6000), listOf(7000,8000,9000), listOf(10000))), 45000)
    }
}