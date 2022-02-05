import { Component } from '@angular/core';
import { UsersService } from './users.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  public data:any = [];
  constructor( private api:UsersService) {
  }
  ngOnInit() {
    this.api.getUsers().subscribe(res => {
      this.data = res;
      console.log(res);
    });
  }
}
