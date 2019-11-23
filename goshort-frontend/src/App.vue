<template>
  <v-app>
    <v-content>
      <URLBox/>
    </v-content>
  </v-app>
</template>

<script>
import URLBox from './components/URLBox';

export default {
  name: 'App',

  components: {
    URLBox,
  },

  data: () => ({
    //
  }),
};
</script>
