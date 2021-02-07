import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

import { UserService } from './user.service';

@Injectable({
	providedIn: 'root'
})
export class AuthService {

	public isLoggedIn: boolean | undefined;

	constructor (
		private httpClient: HttpClient,
		private userService: UserService,
		private router: Router,
	) { }

	public login (payload: string) {
		this.httpClient.post<any>('/api/login', payload, {}).subscribe (
			(resp: any) => {
				this.isLoggedIn = true;
				this.userService.setInfo(resp)
				this.router.navigate(['/first'])
			}
		)
	}

	public logout () {
		this.httpClient.get<any>("/api/logout", {}).subscribe(resp => {
			this.isLoggedIn = false;
		})
	}
}
