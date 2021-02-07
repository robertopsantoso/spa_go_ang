import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MessageService } from '@services/message.service';


@Component({
	selector: 'app-register',
	templateUrl: './register.component.html',
	styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

	user = new FormGroup({
		firstname: new FormControl('', Validators.required),
		lastname: new FormControl('', Validators.required),
		email: new FormControl('', Validators.required),
		password: new FormControl('', Validators.required)
	});

	constructor(
		private httpClient: HttpClient,
		private messageService: MessageService
	) {}

	ngOnInit(): void {
		this.messageService.add("Load register page")
	}

	register() {
		this.messageService.add("Trying to register with data " + JSON.stringify(this.user.value));


		this.httpClient.post<any>('/api/register', JSON.stringify(this.user.value), {}).subscribe (
			(resp: any) => {
				console.log(resp);
			}
		)

		console.log(this.messageService.messages)
	}

}
