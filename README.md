<h1>Xsolla Login Integration with Zendesk Chat</h1>

<p>You can integrate Xsolla Login with Zendesk Chat by linking a user of your Login project to a Zendesk Chat visitor. The external_id parameter that is used at Zendesk’s side corresponds with the user ID in your Login project.</p>

<p>
To link a visitor:
    <ol>
    <li>
        Authenticate a user in your Login project by one of the following requests:
        <ul>
            <li>authentication via username and password (<a href="https://developers.xsolla.com/login-api/jwt/auth-by-username-and-password">JWT</a> and <a href="https://developers.xsolla.com/login-api/oauth-20/oauth-20-auth-by-username-and-password">OAuth 2.0</a>).</li>
            <li>authentication via social networks (<a href="https://developers.xsolla.com/login-api/jwt/jwt-auth-via-social-network/">JWT</a> and <a href="https://developers.xsolla.com/login-api/oauth-20/oauth-20-auth-via-social-network">OAuth 2.0</a>).</li>
        </ul>
    </li>
    <li>Generate a Zendesk JWT. </li>
    <li><a href="https://support.zendesk.com/hc/en-us/articles/360022185314">Authenticate</a> a visitor into Zendesk Chat.</li>
    </ol>
</p>