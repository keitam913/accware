import { useEffect } from 'react';
import uuid from 'uuid';

function getCallPageUrl() {
        const u = new URL(window.location.href);
        u.pathname = '/callback';
        return u;
}

function getAuthUrl() {
        const url = new URL('https://accounts.google.com/o/oauth2/v2/auth');
        url.searchParams.append('redirect_uri', getCallPageUrl().toString());
        url.searchParams.append('scope', 'openid email');
        url.searchParams.append('response_type', 'id_token');
        url.searchParams.append('nonce', uuid.v4());
        url.searchParams.append('client_id', '200301199995-a5uqflvp7gf4ejube19cfnpl45vqrvlj.apps.googleusercontent.com')
        return url;
}

function Login() {
        useEffect(() => {
                window.location = getAuthUrl();
        }, []);
        return null;
}

export default Login;
