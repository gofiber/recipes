<template>
    <div id="user">
        <nav class="text-sm font-semibold mb-6" aria-label="Breadcrumb">
            <ol class="list-none p-0 inline-flex">
                <li class="flex items-center text-blue-500">
                    <a href="#" class="text-gray-700">Home</a>
                    <svg
                            class="fill-current w-3 h-3 mx-3"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 320 512"
                    >
                        <path
                                d="M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 101.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z"
                        />
                    </svg>
                </li>
                <li class="flex items-center">
                    <a href="#" class="text-gray-600">Users</a>
                </li>
            </ol>
        </nav>
        <div class="lg:flex justify-between items-center mb-6">
            <p class="text-2xl font-semibold mb-2 lg:mb-0">Users</p>
        </div>
        <div class="flex flex-wrap -mx-3 mb-20">
            <div class="w-full xl:w-full px-3">
                <div class="w-full bg-white border rounded-lg flex items-center p-6 mb-6 xl:mb-0">
                    <div v-if="loading" class="loading">Loading...</div>
                    <datatable v-else :columns="columns" :data="rows"></datatable>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import {GetUserList} from "../store/api/users";

    export default {
        data: () => ({
            loading: false,
            columns: [
                {label: "#", field: "num", headerClass: "w-1/6"},
                {
                    label: "Name",
                    field: "name"
                },
                {
                    label: "Age",
                    field: "age"
                }
            ],
            rows: []
        }),
        created() {
            this.fetchData();
        },
        watch: {
            // call again the method if the route changes
            $route: "fetchData"
        },
        methods: {
            fetchData() {
                this.loading = true;
                GetUserList()
                    .then(({data}) => {
                        this.loading = false;
                        for (let i in data) {
                            data[i].num = parseInt(i) + 1;
                        }
                        this.rows = data;
                    })
                    .error(err => {
                        this.loading = false;
                    });
            }
        }
    };
</script>
