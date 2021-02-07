import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { tap } from 'rxjs/operators';

import { User } from '@models/user';

@Injectable({
	providedIn: 'root'
})
export class UserService {

	private info: User;

	constructor (
		private httpClient: HttpClient,
	) { }

	public getInfo (): Observable<User> {
		if (this.info) {
			return of(this.info)
		} else {
			return this.reqInfo().pipe(
				tap(resp => this.info = resp)
			)
		}
	}

	public setInfo (user: User) {
		this.info = user
	}

	private reqInfo (): Observable<User> {
		return this.httpClient.get<User>('/api/account', {})
	}
}
