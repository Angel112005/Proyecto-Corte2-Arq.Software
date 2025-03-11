import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
// import { WebSocketService } from './usecases/websocket.service';
import { WebSocketService } from './adapters/websocket.service';



import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { HomeComponent } from './pages/home/home.component';
import { DoctorsComponent } from './pages/doctors/doctors.component';
import { AppointmentsComponent } from './pages/appointments/appointments.component';
import { AdminComponent } from './pages/admin/admin.component';
import { AppointmentFormComponent } from './pages/appointment-form/appointment-form.component';
import { DoctorFormComponent } from './pages/doctor-form/doctor-form.component';
import { FooterComponent } from './components/footer/footer.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    HomeComponent,
    DoctorsComponent,
    AppointmentsComponent,
    AdminComponent,
    AppointmentFormComponent,
    DoctorFormComponent,
    DoctorFormComponent,
    FooterComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,

  ],
  providers: [WebSocketService],
  bootstrap: [AppComponent]
})
export class AppModule { }
