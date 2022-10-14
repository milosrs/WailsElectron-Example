<script lang="ts">
    import {EventsOn} from '../../../wailsjs/runtime/runtime'
    import type { Profile, User, Selection } from "src/types/User";

    let users: User[] = [
        {
            id: 1,
            name: "Dante",
            lastName: "Exum",
        },
        {
            id: 2,
            name: "Zack",
            lastName: "Leday",
        },
        {
            id: 3,
            name: "Nemanja",
            lastName: "Trifunovic",
        },
        {
            id: 4,
            name: "Zeljko",
            lastName: "Obradovic",
        },
        {
            id: 5,
            name: "Ostoja",
            lastName: "Miailovic",
        },
        {
            id: 6,
            name: "Aleksa",
            lastName: "Avramovic",
        },
    ]

    let profiles: Profile[] = [
        {
            id: 1,
            name: "Point Guard",
            tag: "1",
        },
        {
            id: 2,
            name: "Shooting Guard",
            tag: "2",
        },
        {
            id: 3,
            name: "Small Forward",
            tag: "3",
        },
        {
            id: 4,
            name: "Power Forward",
            tag: "4",
        },
        {
            id: 5,
            name: "Center",
            tag: "5",
        },
        {
            id: 6,
            name: "Coach",
            tag: "6",
        },
    ]

    let blacklist: string[] = []

    let selectedUser = 0;
    let selectedProfile = 0;
    let profileAssigns: [User, Profile][] = []

    function selectItem<T extends User | Profile>(item: T, selection: Selection) {
        const items = selection === 'profile' ? profiles : users;
        const activeItem = items.findIndex(x => x.id === item.id);
        const id = items[activeItem].id;

        if(selection === 'profile') {
            selectedProfile = id
        } else {
            selectedUser = id
        }

        console.log(selectedUser, selectedProfile);
    }

    function assign() {
        const pa = profileAssigns.slice();
        pa.push([
            users.find(x => x.id === selectedUser), 
            profiles.find(x => x.id === selectedProfile)
        ]);
        profileAssigns = pa;

        selectedUser = 0;
        selectedProfile = 0;

        console.log(profileAssigns)
    }

    function blacklist_updated(data: any) {
        const bl = blacklist.slice();
        const items = [...data] as string[];
        bl.push(...items)

        blacklist = bl;
    }

    EventsOn('blacklist_updated', blacklist_updated);
</script>

<main id="main" class="overflow-hidden">
<div class="container text-center">
    <h1 class="h3 mb-3 fw-normal">Assign users to profiles</h1>
    <div class="row g-6">
        <div class="col g-4">
            <h4 class="h4 mb-3 fw-normal">Users</h4>
            <table class="table table-dark table-hover">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>First Name</th>
                        <th>Last Name</th>
                    </tr>
                </thead>
                <tbody>
                    {#each users as u}
                        <tr 
                            class="clickable"
                            class:activeRow="{u.id === selectedUser}"
                            on:click={() => selectItem(u, 'user')}
                        >
                            <td>{u.id}</td>
                            <td>{u.name}</td>
                            <td>{u.lastName}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
        <div class="col g-4">
            <h4 class="h4 mb-3 fw-normal">Profiles</h4>
            <table class="table table-dark table-hover">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Profile Name</th>
                        <th>Tag</th>
                    </tr>
                </thead>
                <tbody>
                    {#each profiles as p}
                        <tr 
                            class="clickable"
                            class:activeRow="{p.id === selectedProfile}"
                            on:click={() => selectItem(p, 'profile')}
                        >
                            <td>{p.id}</td>
                            <td>{p.name}</td>
                            <td>{p.tag}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </div>
    <div class="row g-6">
        <div class="col g-6 btn-container">
            <button 
                style="margin-top: 10%;"
                type="button" 
                disabled={selectedUser === 0 || selectedProfile === 0} 
                class="btn btn-primary"
                on:click={assign}>
                Assign selections
            </button>
        </div>
        <div class="col g-6">
            <h4 class="h4 mb-3 fw-normal"> Assignations </h4>
            <table class="table table-dark table-hover">
                <thead>
                    <tr>
                        <th>User Data</th>
                        <th>Profile Name</th>
                    </tr>
                </thead>
                <tbody>
                    {#each profileAssigns as pa}
                        <tr>
                            <td>{pa[0].name + ' ' + pa[0].lastName}</td>
                            <td>{pa[1].name}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </div>
    <div class="row g-6" style="margin-top: 40px;">
        <h4 class="h4 mb-3 fw-normal">Blacklist URLS</h4>
        <ul class="list-group blacklist overflow-auto">
            {#each blacklist as blItem}
                <li class="list-group-item list-group-item-dark blacklistItem">{blItem}</li>
            {/each}
        </ul>
    </div>
</main>

<style>
#main {
    height: 100%;
    width: 100%;
}

.clickable {
    cursor: pointer;
}

.btn-container {
    display: flex;
    justify-content: flex-start;
    align-items: flex-start;
}

.activeRow {
    border-width: 2px;
    border-color: blue;
}

.blacklist {
    margin-top: 40px;
    max-height: 250px;
}

.blacklistItem {
    text-overflow: ellipsis;
}

</style>
  
  