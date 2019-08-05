<template>
<div>
  <p>count: {{ count }}</p>
  <button @click="incr">+</button>
  <button @click="decr">-</button>
</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

import store, { CounterStore } from './store'

@Component({})
export default class App extends Vue {
  counter = CounterStore.CreateProxy(store, CounterStore) as CounterStore

  get count(): number {
    return this.counter.count
  }

  incr(): void {
    this.counter.incr()
  }

  decr(): void {
    this.counter.decr()
  }
}
</script>
