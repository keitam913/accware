import { useEffect } from 'react';
import uuid from 'uuid';

function getCallPageUrl() {
        let u = new URL(window.location.href);
        u.pathname = '/callback';
        return u;
}

function getAuthUrl() {
        let url = new URL(process.env.REACT_APP_AUTH_ENDPOINT);
        url.searchParams.append('client_id', process.env.REACT_APP_CLIENT_ID);
        url.searchParams.append('redirect_uri', getCallPageUrl().toString());
        url.searchParams.append('scope', 'openid email');
        url.searchParams.append('response_type', 'id_token');
        url.searchParams.append('nonce', uuid.v4());
        return url;
}

function Login() {
        useEffect(() => {
                window.location = getAuthUrl();
        }, []);
        return null;
}

export default Login;
