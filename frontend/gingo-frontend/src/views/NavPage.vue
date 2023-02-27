<template >
    <el-menu style="display:flex;justify-content: flex-start; width:100%"
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        @select="handleSelect"
    >
        <el-menu-item index="1">Processing Center</el-menu-item>
        <el-sub-menu index="2">
            <template #title>Workspace</template>
            <el-menu-item index="2-1">item one</el-menu-item>
            <el-menu-item index="2-2">item two</el-menu-item>
            <el-menu-item index="2-3">item three</el-menu-item>
            <el-sub-menu index="2-4">
                <template #title>item four</template>
                <el-menu-item index="2-4-1">item one</el-menu-item>
                <el-menu-item index="2-4-2">item two</el-menu-item>
                <el-menu-item index="2-4-3">item three</el-menu-item>
            </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="3" disabled>Info</el-menu-item>
        <el-menu-item index="4">Orders</el-menu-item>
    </el-menu>
    
    <action-bar style="display:flex;justify-content: flex-end;"></action-bar>
</template>

<script>

import ActionBar from '../views/ActionBar.vue'

export default {
    components:{
        // eslint-disable-next-line vue/no-unused-components
        ActionBar
    }
}
</script>
<style>
.flex-grow {
  flex-grow: 1;
}
</style>