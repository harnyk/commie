#!/usr/bin/env node
import dedent from 'dedent';
import pkg from '../package.json' assert { type: 'json' };
import { fail } from 'node:assert';

const manifest = {
    name: 'executable',
    version: pkg.version,
    description: 'This is a test executable koop',
    prompts: {
        default: {
            selfInvoke: true,
            args: ['-p', 'default'],
        },
    },
    tools: [
        {
            name: 'factorial',
            selfInvoke: true,
            args: ['-t', 'factorial'],
            parameters: {
                type: 'object',
                properties: {
                    n: {
                        type: 'integer',
                        minimum: 0,
                    },
                },
                required: ['n'],
            },
        },
        {
            name: 'power',
            selfInvoke: true,
            args: ['-t', 'power'],
            parameters: {
                type: 'object',
                properties: {
                    a: {
                        type: 'number',
                    },
                    b: {
                        type: 'number',
                    },
                },
                required: ['a', 'b'],
            },
        },
    ],
};

function promptDefault() {
    return dedent`You are a mathematician.
        You can calculate the factorial of a number.`;
}

function toolFactorial(params: { n: number }) {
    const { n } = params;
    let result = 1;
    for (let i = 1; i <= n; i++) {
        result *= i;
    }
    return result;
}

switch (process.argv[2]) {
    case '-p': {
        const promptName = process.argv[3];
        switch (promptName) {
            case 'default': {
                console.log(JSON.stringify(promptDefault()));
                break;
            }
            default: {
                throw new Error(`Unknown prompt: ${promptName}`);
            }
        }
        break;
    }
    case '-t': {
        const toolName = process.argv[3];
        const params = JSON.parse(
            process.env['commie.koop.tool.parameters'] ??
                process.env['commie_koop_tool_parameters'] ??
                fail('missing tool parameters')
        );

        switch (toolName) {
            case 'factorial': {
                console.log(JSON.stringify(toolFactorial(params)));
                break;
            }
            case 'power': {
                console.log(JSON.stringify(params.a ** params.b));
                break;
            }
            default: {
                throw new Error(`Unknown tool: ${toolName}`);
            }
        }
        break;
    }
    default: {
        console.log(JSON.stringify(manifest));
    }
}
