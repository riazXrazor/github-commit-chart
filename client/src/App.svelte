<script>
  import { Highlight } from "svelte-highlight";
  import { markdown } from "svelte-highlight/languages";
  import { anOldHope } from "svelte-highlight/styles";

  export let header;
  let repo = "";
  let url = window.location.pathname;

  const checkRepo = event => {
    const key = event.key;
    const keyCode = event.keyCode;
    if (key === "Enter" || keyCode === 13) {
      if (!repo) {
        alert("Please enter username/repo");
        return;
      }

      fetch(`/api/github/check/${repo}`)
        .then(r => r.json())
        .then(r => {
          if (r.status === 200) {
            window.location.href = `/${repo}`;
          } else {
            alert(`Repo "${repo}" does not exists.`);
          }
        });
    }
  };

  const onCopy = () => {
    navigator.clipboard
      .writeText(code)
      .then(() => alert("successfully copied to clipboard!"))
      .catch(e => alert("Could not copy!!"));
  };

  $: code = `## github commit chart over time

[![github commit chart over time](https://commitchart.xyz${url})](https://commitchart.xyz${url})`;
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

  .repoinput {
    width: 30%;
    margin: 0 auto;
  }
  .repoinput input {
    border-radius: 20px;
    width: 100%;
    text-align: center;
  }

  .repoinput input:hover {
    outline: none;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }

  .copy {
    color: #494949 !important;
    text-transform: capitalize;
    text-decoration: none;
    background: #ffffff;
    padding: 10px;
    border: 1px solid #494949 !important;
    display: inline-block;
    border-radius: 50px;
  }

  .copy:hover {
    cursor: pointer;
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
  {#if url === '/'}
    <p>
      Generate chart by browsing to user/repo, for example,
      <a href="https://commitchart.xyz/riazXrazor/udemy-dl">
        riazXrazor/udemy-dl
      </a>
    </p>
    <div class="repoinput">
      <input
        on:keydown={checkRepo}
        bind:value={repo}
        placeholder="Enter Your username/repo eg. riazXrazor/udemy-dl"
        type="text" />
    </div>
  {/if}
  {#if url !== '/'}
    <div class="output">
      <div class="chart">
        <img src="/chart{url}" alt="" />
      </div>
      <div class="code">
        <Highlight language={markdown} {code} />
      </div>

      <button class="copy" on:click={onCopy}>Copy markdown to clipboard</button>

    </div>
  {/if}
</main>
<footer>
  Made with
  <span>&hearts;</span>
  by Riaz Laskar
</footer>
