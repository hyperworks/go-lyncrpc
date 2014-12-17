package lyncrpc

// REF: http://msdn.microsoft.com/en-us/library/office/microsoft.lync.model.contactavailability_di_3_uc_ocs14mreflyncclnt(v=office.15).aspx

/*	Member name	Description
None	A flag indicating that the contact state is to be reset to the default availability that is calculated by Lync and based on current user activity and calendar state.
Free	A flag indicating that the contact is available.
FreeIdle	idle states are machine state and can not be published as user state.
Busy	A flag indicating that the contact is busy and inactive.
BusyIdle	idle states are machine state and can not be published as user state.
DoNotDisturb	A flag indicating that the contact does not want to be disturbed.
TemporarilyAway	A flag indicating that the contact is temporarily un-alertable.
Away	A flag indicating that the contact cannot be alerted.
Offline	A flag indicating that the contact is not available.
Invalid*/

type Availability string

const (
	AvailabilityNone Availability = "None"
	// TODO:
)
