import type {StorybookConfig} from '@storybook-vue/nuxt';

const config: StorybookConfig = {
    "stories": [
        "../components/**/*.mdx",
        "../components/**/*.stories.@(js|jsx|ts|tsx|mdx)",
        '../components/**/stories/*.stories.@(js|jsx|ts|tsx)'
    ],
    "addons": [
        "@chromatic-com/storybook",
        "@storybook/addon-docs"
    ],
    "framework": {
        "name": "@storybook-vue/nuxt",
        "options": {}
    }
};
export default config;