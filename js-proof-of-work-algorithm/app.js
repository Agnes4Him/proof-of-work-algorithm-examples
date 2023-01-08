const crypto = require('crypto')

//let hash = crypto.getHashes()

let max_nonce = 2 ** 32

function proof_of_work(header, difficulty_bits) {

    let difficulty_target =  2 ** (256 - difficulty_bits)

    for (nonce = 0; nonce <= max_nonce; nonce++) {

        let hash_result = crypto.createHash('sha256').update(String(header) + String(nonce)).digest('hex')

        if (parseInt(hash_result, 16) <= difficulty_target) {
            console.log(`Success with nonce ${nonce}`)
            console.log(`Hash is ${hash_result}`)
            return (hash_result, nonce)
        }
    }
    console.log(`Failed after ${nonce} (max_nonce) tries`)
    return nonce
}

let nonce = 0
let hash_result = ""

for (difficulty_bits = 0; difficulty_bits < 32; difficulty_bits++) {
    let difficulty = 2 ** difficulty_bits
    console.log(`Difficulty: ${difficulty} (${difficulty_bits})`)
    console.log("Searching...")
    // Get start time
    let start_time = Date.now()

    // Create block
    let index_block = "This is an index block" + hash_result

    // Find nonce
    hash_result, nonce = proof_of_work(index_block, difficulty_bits)
    
    // Get end time
    let end_time = Date.now()

    // Get time difference
    let time_diff = end_time - start_time

    console.log(`Elapsed time: ${time_diff} seconds`)

    if (time_diff > 0) {
        let hash_power = parseFloat(nonce)/ time_diff
        console.log(`Hashing power: ${hash_power} hashes per second`)
    }
}
