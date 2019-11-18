<template>
  <v-app id="inspire">
    <v-content>
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
                <v-form>
                  <v-row>
                    <v-col
                      cols="10" 
                      sm="10"
                    >
                      <v-text-field
                        id="url"
                        label="URL"
                        name="url"
                        type="string"
                        single-line=""
                      />
                    </v-col>
                    <v-col 
                      cols="2"
                      sm="2"
                      align-center
                    >
                      <v-btn color="primary">Shorten!</v-btn>
                    </v-col>
                  </v-row>
                </v-form>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
  export default {
    props: {
      source: String,
    },
  }
</script>