import { dirname, join } from 'path'
// @ts-ignore
import type { StorybookConfig } from '@storybook/react-webpack5'

const config: StorybookConfig = {
  stories: ['../src/**/*.stories.mdx', '../src/**/*.stories.@(js|jsx|ts|tsx)'],
  addons: [
    getAbsolutePath('@storybook/addon-links'),
    getAbsolutePath('@storybook/addon-themes'),
    {
      name: getAbsolutePath('@storybook/addon-essentials'),
      options: {
        actions: true,
        backgrounds: true,
        controls: true,
        docs: true, // https://github.com/hipstersmoothie/react-docgen-typescript-plugin/issues/83
        viewport: true,
        toolbars: true,
      },
    },
    getAbsolutePath('@storybook/addon-interactions'),
    getAbsolutePath('@storybook/addon-controls'),
    {
      name: getAbsolutePath('@storybook/addon-styling-webpack'),
      options: {
        postCss: {
          implementation: require.resolve('postcss'),
        },
        rules: [
          {
            test: /\.css$/,
            sideEffects: true,
            use: [
              require.resolve('style-loader'),
              {
                loader: require.resolve('css-loader'),
                options: {
                  // Want to add more CSS Modules options? Read more here: https://github.com/webpack-contrib/css-loader#modules
                  modules: {
                    auto: true,
                  },
                  importLoaders: 1,
                },
              },
              {
                loader: require.resolve('postcss-loader'),
                options: {
                  implementation: require.resolve('postcss'),
                },
              },
            ],
          },
        ],
      },
    },
    getAbsolutePath('@chromatic-com/storybook'),
    getAbsolutePath('@storybook/addon-a11y'),
    getAbsolutePath('@storybook/addon-coverage'),
  ],
  framework: {
    name: getAbsolutePath('@storybook/nextjs'),
    options: {
      useSWC: true,
      fsCache: true,
      lazyCompilation: true,
    },
  },
  features: {},
  docs: {
    autodocs: true,
  },
  typescript: {},
}

export default config

function getAbsolutePath(value: string): any {
  return dirname(require.resolve(join(value, 'package.json')))
}
