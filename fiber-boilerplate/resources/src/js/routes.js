import Dashboard from './components/Dashboard'
import DashboardHome from './pages/Home'
import UserPage from './pages/User'

const routes = [
    {path: '/', redirect: {name: 'DashboardHome'}},
    {
        path: '/dashboard', component: Dashboard, children: [
            {path: '/', redirect: {name: 'DashboardHome'}},
            {path: 'home', name: 'DashboardHome', component: DashboardHome},
            {path: 'user', name: 'UserPage', component: UserPage}
        ]
    }
]

export default routes