#!/usr/bin/env node

import dedent from 'dedent';
import { Koop, Tool } from 'commie-koop';

const mathjs = new Tool()
    .name('mathjs')
    .description(
        dedent`
        Evaluate math expressions, using mathjs API.
        Example:
        {
            "expr": [
                "a = 1.2 * (2 + 4.5)",
                "a / 2",
                "5.08 cm in inch",
                "sin(45 deg) ^ 2",
                "9 / 3 + 2i",
                "b = [-1, 2; 3, 1]",
                "det(b)"
            ],
            "precision": 14
        }
        
        NOTE: If you use variables in your expressions, they are only available within a single call to this tool.
        LIMITATION: this tool does not solve equations for you. It only performs calculations.
        `
    )
    .parameters({
        type: 'object',
        properties: {
            expr: {
                type: 'array',
                description: 'The math expressions to evaluate',
                items: {
                    type: 'string',
                },
            },
            precision: {
                type: 'number',
                description: 'The number of decimal places to round to',
                minimum: 0,
            },
        },
        required: ['expr', 'precision'],
    })
    .handler(
        async (params: {
            expr: string;
            precision: number;
        }): Promise<string> => {
            const response = await fetch('https://api.mathjs.org/v4/', {
                method: 'POST',
                body: JSON.stringify({
                    expr: params.expr,
                    precision: params.precision,
                }),
            });
            return await response.json();
        }
    );

const fac = new Tool()
    .name('factorial')
    .parameters({
        type: 'object',
        properties: {
            n: {
                type: 'integer',
                minimum: 0,
            },
        },
        required: ['n'],
    })
    .handler((params: { n: number }) => {
        const { n } = params;
        let result = 1;
        for (let i = 1; i <= n; i++) {
            result *= i;
        }
        return result;
    });

const pow = new Tool()
    .name('power')
    .parameters({
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
    })
    .handler((params: { a: number; b: number }) => params.a ** params.b);

// ----------------

new Koop()
    .name('koop-example-executable')
    .version('1.0.0')
    .description('This is a test executable koop')
    .prompt(
        'default',
        () => dedent`You are a matematician.
        You can calculate factorial, power and other math operations.
        You should not calculate in your mind, always use tools.
        Prefer one-shot solutiuons.
        Do not print intermediate steps.
        IMPORTANT: Output format - plain text. Do not use LaTeX of Mathjax.
        `
    )
    .tool(fac)
    .tool(pow)
    .tool(mathjs)
    .run();
