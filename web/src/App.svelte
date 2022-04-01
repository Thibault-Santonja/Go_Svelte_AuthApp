<script lang="ts">
  import Footer from "./lib/Footer.svelte";
  import "./app.css";
  import Header from "./lib/Header.svelte";
  import Home from "./lib/Home.svelte";
  import Notes from "./lib/Notes.svelte";
</script>


<slot />

<body class="w-full flex flex-col min-h-screen">
  <Header />

  <main class="flex-grow grid justify-items-center gap-8">
    <Home />
    <Notes />
  </main>

  <Footer />
</body>


<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  }

  /* color: #ff3e00; */
</style>
