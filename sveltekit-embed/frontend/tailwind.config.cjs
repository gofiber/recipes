module.exports = {
  mode: "jit",
  content: ["./src/**/*.{html,js,svelte,ts}"],
  darkMode: "class", // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        primary: "#00acd7",
      },
    },
  },
  variants: {},
  plugins: [],
};
