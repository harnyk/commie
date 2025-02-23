const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false,
});

rl.on('line', (line) => {
    const { i: id, p: body } = JSON.parse(line);

    console.error('req', line)

    const payload = JSON.parse(body);

    const { type } = payload;

    switch (type) {
        case 'add': {
            const { a, b } = payload;
            respond(id, { c: a + b });
            break;
        }
        case 'exit':
            respond(id, { msg: 'bye' });
            process.exit(0);
        default:
            break;
    }
});

rl.on('close', () => {
    process.exit(0);
});

function respond(id, payload) {
    const res = JSON.stringify({
        i: id,
        t: 'o',
        p: JSON.stringify(payload),
    })
    console.error('res', res)
    console.log(res);
}