/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'secondary-font': ['"Clicker Script"', 'cursive'],
      }
    },
  },
  daisyui: {
    themes: [
      {
        light: {
          "primary": "#151515", // veryDarkGray
          "secondary": "#222222", // darkGray
          "accent": "#147B74", // themeColor
          "neutral": "#f8fffe", // veryLightGray
          "base-100": "#e5e0df", // lightGray
        },
        dark: {
          "primary": "#e5e0df", // veryDarkGray
          "secondary": "#151515", // darkGray
          "accent": "#147B74", // themeColor
          "neutral": "#f8fffe", // veryLightGray
          "base-100": "#222222", // lightGray df4996
        }
      },
    ],
  },
  plugins: [require("daisyui")],
}
