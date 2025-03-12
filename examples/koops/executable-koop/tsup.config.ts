import { defineConfig } from 'tsup';

export default defineConfig({
    entry: ['bin/main.ts'],
    format: ['esm'],
    dts: false,
    sourcemap: true,
    clean: true,
    outDir: 'lib',
    shims: false,
});
