export enum PageRoutes {
    Index = '/',
    Search = '/search',
    SearchConsole = '/search/console',
    SignIn = '/sign-in',
    SignUp = '/sign-up',
    PostSignUp = '/post-sign-up',
    UnlockAccount = '/unlock-account/:token',
    Welcome = '/welcome',
    Settings = '/settings',
    User = '/user/*',
    Organizations = '/organizations/*',
    SiteAdmin = '/site-admin/*',
    LicenseManagement = '/license-admin/*',
    SiteAdminInit = '/site-admin/init',
    PasswordReset = '/password-reset',
    ApiConsole = '/api/console',
    UserArea = '/users/:username/*',
    Survey = '/survey/:score?',
    Extensions = '/extensions',
    Help = '/help/*',
    Debug = '/-/debug/*',
    RepoContainer = '/*',
    SetupWizard = '/setup',
    Teams = '/teams/*',
    RequestAccess = '/request-access/*',
    GetCody = '/get-cody',
    BatchChanges = '/batch-changes/*',
    CodeMonitoring = '/code-monitoring/*',
    Insights = '/insights/*',
    SearchJobs = '/search-jobs/*',
    Contexts = '/contexts',
    CreateContext = '/contexts/new',
    EditContext = '/contexts/:specOrOrg/:spec?/edit',
    Context = '/contexts/:specOrOrg/:spec?',
    NotebookCreate = '/notebooks/new',
    Notebook = '/notebooks/:id',
    Notebooks = '/notebooks',
    SearchNotebook = '/search/notebook',
    CodySearch = '/search/cody',
    Cody = '/cody/chat',
    Own = '/own',
    AppAuthCallback = '/app/auth/callback',
    AppSetup = '/app-setup',
}
