import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

export default defineConfig({
    build:{
        assetsDir: 'assets',
    },
    plugins: [
        vue(),
        AutoImport({
            resolvers: [ElementPlusResolver()],
        }),
        Components({
            resolvers: [ElementPlusResolver()],
        }),
    ],
    server: {
        proxy: {
            '/login': {
                "target": "http://localhost:8080",
            },
            '/register': {
                "target": "http://localhost:8080",
            },
            '/admin/upload': {
                "target": "http://localhost:8080",
            }
        }
    }
})