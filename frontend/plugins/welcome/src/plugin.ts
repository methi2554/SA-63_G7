 
import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Drug from './components/Disease'
import Login from './components/Login'
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/disease', Drug);
    router.registerRoute('/welcome', WelcomePage);
    

  },
});
 
