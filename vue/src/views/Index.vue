<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="evan-forbes.stevia.stevia" type="sweetner" :fields="[ ['creator', 1, 'string'] , ['calories', 2, 'int32'] , ['organic', 3, 'bool'] , ['source', 4, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
