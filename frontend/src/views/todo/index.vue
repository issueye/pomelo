<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { reactive, ref } from 'vue';

const activeName = ref('')
const itemVisible = ref(false)
const itemForm = reactive({
  id: 0,
  name: '',
})

const onAddItemClick = () => {
  itemVisible.value = true;
}


</script>

<template>
  <div class="flex flex-row bg-white h-full w-full">
    <div class="flex flex-col justify-between border-r border-b-[#D9D9D9] w-40 p-2">
      <el-collapse v-model="activeName" accordion>
        <el-collapse-item name="1">
          <template #title>
            <Icon icon="material-symbols:menu-rounded" /> <span class="ml-3">工作</span>
          </template>
          测试内容
        </el-collapse-item>
        <el-collapse-item title="学习" name="2">
          <template #title>
            <Icon icon="material-symbols:menu-rounded" /> <span class="ml-3">学习</span>
          </template>
          测试内容
        </el-collapse-item>
      </el-collapse>

      <div class="flex justify-center">
        <el-button type="primary" circle class="shadow-lg mb-2" @click="onAddItemClick">
          <Icon icon="ic:baseline-plus" />
        </el-button>
      </div>
    </div>
    <div class="flex flex-grow">
      main
    </div>
  </div>

  <el-dialog v-model="itemVisible" title="添加类型" draggable :close-on-click-modal="false">
    <el-form :model="itemForm" label-width="auto">
      <el-form-item label="类型名称" prop="name">
        <el-input v-model="itemForm.name" />
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
