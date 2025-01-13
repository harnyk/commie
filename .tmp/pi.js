const { performance } = require('perf_hooks');
const start = performance.now();

let i = 0;
let n = 1e6;
let sum = 0;

while(i++ < n){
    sum += Math.sqrt(Math.PI);
}

const end = performance.now();

console.log(`It took ${end - start} milliseconds to compute the square root of PI ${n} times.`);
