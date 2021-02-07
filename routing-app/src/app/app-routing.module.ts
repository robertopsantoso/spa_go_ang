import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// Services
import { AuthGuard } from '@helpers/auth.guard';

// Components
import { FirstComponentComponent } from './first-component/first-component.component';
import { ChildAComponent } from './first-component/child-a/child-a.component';
import { ChildBComponent } from './first-component/child-b/child-b.component';
import { SecondComponentComponent } from './second-component/second-component.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

const routes: Routes = [
	{path: 'login', component: LoginComponent, canActivate: [AuthGuard]},
	{path: 'register', component: RegisterComponent, canActivate: [AuthGuard]},
	{
		path: 'first',
	 	component: FirstComponentComponent,
	 	children: [
	 		{
	 			path: 'child-a',
	 			component: ChildAComponent
	 		},
	 		{
	 			path: 'child-b',
	 			component: ChildBComponent
	 		}
	 	],
	    canActivate: [AuthGuard]
	},
	{path: 'second', component: SecondComponentComponent},
	{
	    path: 'profile',
	    loadChildren: () => import('./profile/profile.module').then(m => m.ProfileModule),
	    canLoad: [AuthGuard]
	},
	{
	    path: 'task',
	    loadChildren: () => import('./task/task.module').then(m => m.TaskModule),
	    canLoad: [AuthGuard]
	},
	{path: '', redirectTo: '/first', pathMatch: 'full'}, //redirect to 'first component'
	{path: '**', component: PageNotFoundComponent}, //Wildcard route for a 404 page
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
