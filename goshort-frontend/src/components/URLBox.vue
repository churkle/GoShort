<template>
  <v-container
    class="fill-height"
    fluid
  >
      <v-row
        align="center"
        justify="center"
      >
        <v-col
          cols="12"
          sm="8"
          md="6"
        >
          <v-card class="elevation-12">
            <v-card-text>
              <transition name="fade">
                <v-row v-if="showForm">
                  <v-col
                    cols="10" 
                    sm="10"
                  >
                    <v-text-field
                      v-model="url"
                      id="urlField"
                      label="URL"
                      name="urlField"
                      type="string"
                      placeholder="www.example.com"
                    />
                  </v-col>
                  <v-col 
                    cols="2"
                    sm="2"
                    align-center
                  >
                    <v-btn color="primary" v-on:click="submitUrl">Shorten!</v-btn>
                  </v-col>
                </v-row>
              </transition>
              <v-row v-if="!showForm">
                <v-col
                    cols="10" 
                    sm="10"
                >
                  <v-text-field
                      id="shortLinkBox"
                      label="Your Shortened Link!"
                      name="shortLinkBox"
                      :hint=copiedHint
                      persistent-hint="true"
                      type="string"
                      readonly="true"
                      v-model="shortLink"
                      outlined="true"
                      rounded="true"
                  />
                </v-col>
                <v-col 
                    cols="2"
                    sm="2"
                    align-center
                >
                  <v-btn color="primary"
                  v-clipboard:copy="shortLink"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError">Copy</v-btn>
                </v-col>
              </v-row>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
import axios from "axios";

export default {
  name: 'URLBox',

  data: () => ({
    url: '',
    shortLink: '',
    showForm: 'true',
    copiedHint: '',
  }),
  methods: {
    submitUrl: function () {
      axios.post("http://localhost:8080/api/v1/shorten",{
        url: this.url
      }).then((response) => {
        console.log(response)
        this.shortLink = "http://localhost:3030/" + response.data
        console.log(this.shortLink)
        this.showForm = ''
      });
    },
    onCopy: function () {
      this.copiedHint = 'Link copied to clipboard!'
    }
  }
};
</script>
