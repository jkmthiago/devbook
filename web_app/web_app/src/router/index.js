import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import Home from '@/views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '',
      component: Login
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
      meta: { requiresAuth: true }
    },
  ]
})

function getCookie(nome) {
  let nomeCookie = nome + "=";
  let ca = document.cookie.split(';');
  for(var i = 0;  i < ca.length; i++){
      var c = ca[i];
      while (c.charAt(0) == ' ') c = c.substring(1, c.length);
      if (c.indexOf(nomeCookie) == 0) return c.substring(nomeCookie.length, c.length);
  }
  return null;
}

function isAuthenticated() {
  const token = getCookie('Token');
  return !!token;
}

// Função para verificar necessidade de autenticação de acesso para a rota
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!isAuthenticated()) {
      next({ // Em caso de não estar autenticado redireciona a tela de login
        path: '/login',
        query: { redirect: to.fullPath }
      });
    } else {
      next(); // Em caso do usuário estive autenticado
    }
  } else {
    next(); // Em caso de não necessidade de autenticação para acesso a rota
  }
})

export default router
