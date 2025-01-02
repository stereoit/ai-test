/* eslint-env node */
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,js,svelte,ts}",
    "./node_modules/@smui/**/*.svelte",
    "./node_modules/svelte-material-ui/**/*.svelte", 
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

