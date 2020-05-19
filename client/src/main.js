import App from "./App.svelte";

const app = new App({
  target: document.body,
  props: {
    header: "github commit chart",
  },
});

export default app;
