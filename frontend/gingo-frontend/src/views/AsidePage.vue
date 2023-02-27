// eslint-disable-next-line vue/valid-template-root
<template>
    <el-tree :data="treeData" :props="defaultProps" @node-click="handleNodeClick" />
</template>
  
<script>

export default {
  data() {
    return {
        defaultProps: {
            children: 'children',
            label: 'label',
        },
        treeData: [
            {
            label: 'Level one 1',
            children: [
                {
                label: 'Level two 1-1',
                children: [
                    {
                    label: 'Level three 1-1-1',
                    },
                ],
                },
            ],
            },
            {
            label: 'Level one 2',
            children: [
                {
                label: 'Level two 2-1',
                children: [
                    {
                    label: 'Level three 2-1-1',
                    },
                ],
                },
                {
                label: 'Level two 2-2',
                children: [
                    {
                    label: 'Level three 2-2-1',
                    },
                ],
                },
            ],
            },
            {
            label: 'Level one 3',
            children: [
                {
                label: 'Level two 3-1',
                children: [
                    {
                    label: 'Level three 3-1-1',
                    },
                ],
                },
                {
                label: 'Level two 3-2',
                children: [
                    {
                    label: 'Level three 3-2-1',
                    },
                ],
                },
            ],
            },
        ]
    }
  },
  methods: {
    handleNodeClick (data) {
        console.log(data)
    }
  },
}
</script>
