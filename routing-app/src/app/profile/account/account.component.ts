import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'

import { User } from '@models/user';
import { UserService } from '@services/user.service';

@Component({
	selector: 'app-account',
	templateUrl: './account.component.html',
	styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

	info: User;

	constructor (
		private userService: UserService,		
	) { }

	ngOnInit (): void {
		this.userService.getInfo().subscribe(resp => {this.info = resp})
	}

}
