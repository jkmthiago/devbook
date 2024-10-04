<script setup>

import { RouterView, useRoute } from 'vue-router'
import Footer from './components/Footer.vue';
import Header from './components/Header.vue';
import { computed } from 'vue';

const route = useRoute();

const areWeInTheLoginPage = computed (() => route.name === "login" || route.name === "");
const areWeInTheRegisterPage = computed (() => route.name === "register");

</script>

<template>
  <Header v-if="!areWeInTheLoginPage && !areWeInTheRegisterPage" />
  <RouterView />
  <Footer v-if="!areWeInTheLoginPage && !areWeInTheRegisterPage"/>
</template>

<style>

</style>
