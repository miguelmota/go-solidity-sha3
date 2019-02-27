const ethers = require('ethers')

//console.log(ethers.utils.solidityKeccak256([ 'uint8[]' ], [ [1,2,3] ] ))
//console.log(ethers.utils.solidityKeccak256([ 'bool[]' ], [ [true, false, true] ] ))
//console.log(ethers.utils.solidityKeccak256([ 'int8[]' ], [ [1,2,3] ] ))

//console.log(ethers.utils.solidityKeccak256([ 'string[]' ], [ ['a', 'b', 'c'] ] ))

//console.log(ethers.utils.solidityKeccak256([ 'address[]' ], [ ['0xa', '0xb', '0xc'] ] ))

//console.log(ethers.utils.solidityKeccak256([ 'address' ], [ '0xa'] ))
//console.log(ethers.utils.solidityKeccak256([ 'address' ], [ '0x0'] ))


//console.log(ethers.utils.solidityKeccak256([ 'bytes32' ], [ '0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45' ] ))
//console.log(ethers.utils.solidityKeccak256([ 'bytes1' ], [ '0x4' ] ))
//console.log(ethers.utils.solidityKeccak256(['bytes32'],['0x0a00000000000000000000000000000000000000000000000000000000000000']))

//var msgDigest = ethers.utils.keccak256(Buffer.from('abc'));

//console.log(msgDigest.toString())

//console.log(ethers.utils.solidityKeccak256(['address', 'uint256'],['0x935F7770265D0797B621c49A5215849c333Cc3ce', '100000000000000000']))

/*
console.log(ethers.utils.solidityKeccak256([
  'address', 'bytes1', 'uint8[]', 'bytes32', 'uint256', 'address[]', 'uint32' ],
  [
    '0x935F7770265D0797B621c49A5215849c333Cc3ce',
    '0xa',
    [128, 255],
    '0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45',
    '100000000000000000',
    [
      '0x13D94859b23AF5F610aEfC2Ae5254D4D7E3F191a',
      '0x473029549e9d898142a169d7234c59068EDcBB33'
    ],
    123456789
  ]
))
// 0xad390a98c1c32cdb1f046f6887a4109f12290b690127e6e15da4ca210235510e
*/
