<script>
  import { Highlight } from "svelte-highlight";
  import { markdown } from "svelte-highlight/languages";
  import { anOldHope } from "svelte-highlight/styles";
  export let header;

  let url = window.location.pathname;

  // fetch(`/chart${url}`)
  //   .then(r => r.text())
  //   .then(r => console.log(r));

  $: code = `## github commit chart over time

[![github commit chart over time](https://commitchart.cc${url})](https://commitchart.cc${url})`;
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  .chart,
  .code {
    width: 50%;
    margin: 0 auto;
  }

  .code {
    text-align: left;
    font-size: 1.1em;
  }

  .chart img {
    width: 100%;
    height: auto;
  }

  h1 {
    color: #151516;
    text-transform: uppercase;
    font-size: 2em;
    font-weight: bolder;
  }

  h1 small {
    display: block;
    font-size: 0.3em;
    font-weight: normal;
  }

  footer {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    text-align: center;
    padding: 20px 0;
  }

  footer span {
    color: rgba(255, 0, 0, 0.452);
    font-size: 1.2em;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<svelte:head>
  {@html anOldHope}
</svelte:head>
<main>

  <h1>
    {header}
    <small>Plot github repository commit over time.</small>
  </h1>
  <p>
    Generate chart by browsing to user/repo, for example,
    <a href="riazXrazor/udemy-dl">riazXrazor/udemy-dl</a>
  </p>
  {#if url.length}
    <div class="output">
      <div class="chart">
        <img src="/chart{url}" alt="" />
      </div>
      <div class="code">
        <Highlight language={markdown} {code} />
      </div>
    </div>
  {/if}
</main>
<footer>
  Made with
  <span>&hearts;</span>
  by Riaz Laskar
</footer>
