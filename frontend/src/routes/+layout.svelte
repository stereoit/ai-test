

<script lang="ts">
  import { onMount } from 'svelte';
  import TopAppBar, {
    Row,
    Section,
    Title,
    AutoAdjust,
  } from '@smui/top-app-bar';
  import IconButton, {Icon} from '@smui/icon-button';
  import {mdiAccount, mdiWeatherSunny, mdiWeatherNight} from '@mdi/js';
  import {Svg} from '@smui/common';
  import { mdiMenu } from '@mdi/js';
  import Drawer from '$lib/Drawer.svelte';

  let { children } = $props();
  let topAppBar: TopAppBar | null = $state(null);
  let darkTheme: boolean | undefined = $state(false);
  let isMenuOpen: boolean = $state(false);

  onMount(() => {
    darkTheme = window.matchMedia('(prefers-color-scheme: dark)').matches;
  })
</script>


<svelte:head>
    <!-- SMUI Styles -->
    {#if darkTheme === undefined}
        <link rel="stylesheet" href="/smui.css" media="(prefers-color-scheme: light)" />
        <link rel="stylesheet" href="/smui-dark.css" media="screen and (prefers-color-scheme: dark)" />
    {:else if darkTheme}
        <link rel="stylesheet" href="/smui.css" media="(prefers-color-scheme: light)" />
	    <link rel="stylesheet" href="/smui-dark.css" media="screen" />
    {:else}
        <link rel="stylesheet" href="/smui.css"/>
    {/if}
    <link rel="stylesheet" href="/app.css" />
</svelte:head>

<TopAppBar bind:this={topAppBar} variant="standard">
  <Row>
    <Section onclick={() => isMenuOpen = !isMenuOpen}>
      <IconButton><Icon tag="svg" viewBox="0 0 24 24"><path fill="currentColor" d={mdiMenu} /></Icon></IconButton>
      <Title>Phones</Title>
    </Section>
    <Section align="end" toolbar>
      <IconButton onclick={() => darkTheme = !darkTheme} title={darkTheme ? 'Lights on.' : 'Lights off.'}>
        <Icon tag="svg" viewBox="0 0 24 24">
          <path fill="currentColor" d={darkTheme ? mdiWeatherNight : mdiWeatherSunny} />
        </Icon>
      </IconButton>
      
    </Section>
  </Row>
</TopAppBar>

<Drawer open={isMenuOpen}/>

<AutoAdjust {topAppBar}>
  {@render children?.()}
</AutoAdjust>


