import Vue from 'vue'
import Vuex from 'vuex'
import { Module, VuexModule, getter, mutation } from 'vuex-class-component'

Vue.use(Vuex)

@Module({ namespacedPath: 'counter/' })
export class CounterStore extends VuexModule {
  @getter
  public count: number = 0

  @mutation
  incr(): void {
    this.count++
  }

  @mutation
  decr(): void {
    this.count--
  }
}

export default new Vuex.Store({
  modules: {
    counter: CounterStore.ExtractVuexModule(CounterStore)
  }
})
