import obfuscatorPlugin from "vite-plugin-javascript-obfuscator";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [
    obfuscatorPlugin({
      options: {
        numbersToExpressions: true,
        debugProtection: true,
        debugProtectionInterval: 1000,
      },
    }),
  ],
});